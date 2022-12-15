package localService

import (
	"bytes"
	"context"
	"go2ban/cmd/firewall"
	"go2ban/pkg/config"
	"go2ban/pkg/syncMap"
	"go2ban/pkg/validator"
	"log"
	"os"
	"time"
)

func checkLogAndBlock(ctx context.Context, service config.Service, countFailsMap syncMap.SyncMap, endBytesMap syncMap.SyncMap) {
	//TODO ReadAt //find byte //sync map name+date = last byte
	file, err := os.Open(service.LogFile)
	f, err := file.Stat()
	if err != nil {
		log.Println("Local service, can't open log file ", err)
		return
	}
	defer file.Close()
	var startByte int64

	endByte := endBytesMap.Load(service.LogFile)

	if endByte <= f.Size() {
		startByte = endByte
	}
	buf := make([]byte, f.Size()-startByte)

	readB, err := file.ReadAt(buf, startByte)
	if err != nil {
		log.Println("Local service, can't readAt log file ", err)
		return
	}

	endBytesMap.Save(service.LogFile, endByte+int64(readB))

	log.Printf("Bytes read %d of filesize %d\n", readB, f.Size()) //TODO del

	if bytes.ContainsAny(buf, service.Regxp) {
		for bySt := range bytes.Split(buf, []byte{'\n'}) {

			ip, err := validator.CheckIp(string(bySt))

			if err == nil {
				countFailsMap.Increment(ip)
				if int(countFailsMap.Load(ip)) > config.Get().ServiceFails {

					go firewall.BlockIP(ctx, ip)

					log.Printf("Block localservice: %s ip: %s", service.Name, ip)
				}
			}
		}
	}

	time.Sleep(time.Second)
	checkLogAndBlock(ctx, service, countFailsMap, endBytesMap)
}
