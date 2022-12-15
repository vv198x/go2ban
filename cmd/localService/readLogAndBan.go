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
)

func checkLogAndBlock(ctx context.Context, service config.Service, countFailsMap, endBytesMap syncMap.SyncMap) {
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
	} else {
		endBytesMap.Save(service.LogFile, 0)
	}

	buf := make([]byte, f.Size()-startByte)

	readB, err := file.ReadAt(buf, startByte)
	if err != nil {
		log.Println("Local service, can't readAt log file ", err)
		return
	}

	endBytesMap.Save(service.LogFile, endByte+int64(readB))

	log.Printf("Bytes read %d of filesize %d\n", readB, f.Size()) //TODO del

	findBytes := []byte(service.Regxp)
	if bytes.Contains(buf, findBytes) {
		for _, bySt := range bytes.Split(buf, []byte{'\n'}) {
			if !bytes.Contains(bySt, findBytes) {
				continue
			}
			ip, err := validator.CheckIp(string(bySt))

			if err == nil {
				countFailsMap.Increment(ip)
				if int(countFailsMap.Load(ip))+1 == config.Get().ServiceFails {
					go firewall.BlockIP(ctx, ip)

					log.Printf("Block localservice: %s ip: %s", service.Name, ip)
				}
			}
		}
	}
}
