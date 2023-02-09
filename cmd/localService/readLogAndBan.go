package localService

import (
	"bytes"
	"context"
	"github.com/vv198x/go2ban/cmd/firewall"
	"github.com/vv198x/go2ban/pkg/config"
	"github.com/vv198x/go2ban/pkg/storage"
	"github.com/vv198x/go2ban/pkg/validator"
	"log"
	"os"
)

func checkLogAndBlock(ctx context.Context, service config.Service, countFailsMap, endBytesMap storage.SyncMap) {
	file, errO := os.Open(service.LogFile)
	f, err := file.Stat()
	if (err != nil) && (errO != nil) {
		log.Println("Local service, can't open log file ", service.LogFile, err)
		return
	}
	defer file.Close()

	//To start reading
	var startByte int64

	//Keep last file size - by log file and service name
	key := service.Name + service.LogFile

	endByte := endBytesMap.Load(key)

	//If the file has become smaller, read again
	if endByte <= f.Size() {
		startByte = endByte
	} else {
		endBytesMap.Save(key, 0)
	}

	// Do not read 0 bytes
	if f.Size()-startByte == 0 {
		return
	}

	//Read Buffer
	buf := make([]byte, f.Size()-startByte)

	//Read where we finished last
	readB, err := file.ReadAt(buf, startByte)
	if err != nil {
		log.Println("Local service, can't readAt log file ", err)
		return
	}

	findBytes := []byte(service.Regxp)

	for _, bySt := range bytes.Split(buf, []byte{'\n'}) {
		if !bytes.Contains(bySt, findBytes) {
			continue
		}
		ip, err := validator.CheckIp(string(bySt))

		if err == nil {
			countFailsMap.Increment(ip)
			count := int(countFailsMap.Load(ip))

			if count == config.Get().ServiceFails {

				go firewall.BlockIP(ctx, ip)

				log.Printf("Block localservice: %s ip: %s", service.Name, ip)
			}
		}
	}

	endBytesMap.Save(key, endByte+int64(readB))

	//log.Printf("Bytes read %d of filesize %d, file: %s, \non seconds %.4f",
	//	readB, f.Size(), f.Name(), time.Since(start).Seconds())
}
