package localService

import (
	"bytes"
	"context"
	"go2ban/cmd/firewall"
	"go2ban/pkg/config"
	"go2ban/pkg/storage"
	"go2ban/pkg/validator"
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

	//Для начала чтения
	var startByte int64

	//Сохранять последний размер файла - по лог файлу и имени сервиса
	key := service.Name + service.LogFile

	endByte := endBytesMap.Load(key)

	//Если файл стал меньше, читаем заново
	if endByte <= f.Size() {
		startByte = endByte
	} else {
		endBytesMap.Save(key, 0)
	}

	// Не читаем 0 байт
	if f.Size()-startByte == 0 {
		return
	}

	//Буфер чтения
	buf := make([]byte, f.Size()-startByte)

	//Читаем откуда закончили в последний раз
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
