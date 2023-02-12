package localService

import (
	"bytes"
	"context"
	"github.com/vv198x/go2ban/cmd/firewall"
	"github.com/vv198x/go2ban/cmd/validator"
	"github.com/vv198x/go2ban/config"
	"github.com/vv198x/go2ban/pkg/storage"
	"log"
	"os"
)

func (s *serviceWork) checkLogAndBlock(ctx context.Context, logFile string, countFailsMap, endBytesMap storage.SyncMap) {
	file, errO := os.Open(logFile)
	f, err := file.Stat()
	if (err != nil) && (errO != nil) {
		log.Println("Local service, can't open log file ", logFile, err)
		return
	}
	defer file.Close() //nolint

	//To start reading
	var startByte int64

	//Keep last file size, service name + file
	key := s.Name + logFile

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

	// String in file
	for _, bySt := range bytes.Split(buf, []byte{'\n'}) {
		// String for find
		for _, findBytes := range s.FindSt {

			if !bytes.Contains(bySt, findBytes) {
				continue
			}
			ip, err := validator.CheckIp(string(bySt))

			if err == nil {
				countFailsMap.Increment(ip)
				count := int(countFailsMap.Load(ip))

				if count == config.Get().ServiceFails {

					go firewall.Do().Block(ctx, ip)

					log.Printf("Block localservice: %s ip: %s", s.Name, ip)
				}
			}
		}
	}

	endBytesMap.Save(key, endByte+int64(readB))
}
