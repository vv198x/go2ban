package localService

import (
	"bytes"
	"context"
	"fmt"
	"go2ban/cmd/firewall"
	"go2ban/pkg/config"
	"go2ban/pkg/storage"
	"go2ban/pkg/validator"
	"log"
	"os"
	"time"
)

func checkLogAndBlock(ctx context.Context, service config.Service, countFailsMap, endBytesMap storage.SyncMap) {
	file, err := os.Open(service.LogFile)
	f, err := file.Stat()
	if err != nil {
		log.Println("Local service, can't open log file ", err)
		return
	}
	defer file.Close()
	var startByte int64
	start := time.Now() //todo del
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

	findBytes := []byte(service.Regxp)

	for _, bySt := range bytes.Split(buf, []byte{'\n'}) {
		if !bytes.Contains(bySt, findBytes) {
			continue
		}

		ip, err := validator.CheckIp(string(bySt))

		if err == nil {
			countFailsMap.Increment(ip)
			count := int(countFailsMap.Load(ip))
			if count == 4 {
				fmt.Println(ip)
			}
			if count == config.Get().ServiceFails {

				go firewall.BlockIP(ctx, ip)

				//fmt.Println(ip, count)
				log.Printf("Block localservice: %s ip: %s", service.Name, ip)
			}
		}
	}

	log.Printf("Bytes read %d of filesize %d, file: %s\n, on second %.4f", readB, f.Size(), f.Name(),
		time.Since(start).Seconds()) //todo del
}
