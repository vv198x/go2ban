package localService

import (
	"bytes"
	"context"
	"errors"
	"github.com/vv198x/go2ban/cmd/firewall"
	"github.com/vv198x/go2ban/cmd/validator"
	"github.com/vv198x/go2ban/config"
	"github.com/vv198x/go2ban/storage"
	"io"
	"log"
	"os"
)

// maxReadChunk caps how many bytes checkLogAndBlock reads in a single
// pass, so a huge backlog on one file (e.g. an unrotated docker json-log
// that grew for months) can't spike memory by gigabytes in one shot.
const maxReadChunk = 64 * 1024 * 1024 // 64MB

func (s serviceWork) checkLogAndBlock(ctx context.Context, logFile string, countFailsMap, endBytesMap storage.Storage) {
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
		startByte = 0
		endBytesMap.Save(key, 0)
	}

	// Do not read 0 bytes
	toRead := f.Size() - startByte
	if toRead == 0 {
		return
	}

	// Cap how much we read in a single pass so an unrotated log that grew
	// unbounded (e.g. months of docker json-log output) can't blow up
	// memory in one allocation; the remainder is picked up next cycle.
	if toRead > maxReadChunk {
		toRead = maxReadChunk
	}

	//Read Buffer
	buf := make([]byte, toRead)

	//Read where we finished last
	readB, err := file.ReadAt(buf, startByte)
	if err != nil && !errors.Is(err, io.EOF) {
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

	endBytesMap.Save(key, startByte+int64(readB))
}
