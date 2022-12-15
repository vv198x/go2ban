package localService

import (
	"context"
	"go2ban/pkg/config"
	"go2ban/pkg/syncMap"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

const nameMapFile = "endBytesMap"

func WorkerStart(services []config.Service) {
	if config.Get().Flags.RunAsDaemon == false {
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	countFailsMap := syncMap.NewCountersMap()
	endBytesMap := syncMap.NewStorageMap()

	saveMapFile := filepath.Join(config.Get().LogDir, nameMapFile)
	err := endBytesMap.ReadFromFile(saveMapFile)
	if err != nil {
		log.Println("Can't read from file map: ", err)
	}

	go func(sleepMinutes time.Duration) {
		for {
			for _, service := range services {
				if service.On {

					go checkLogAndBlock(ctx, service, countFailsMap, endBytesMap)
				}
			}
			time.Sleep(sleepMinutes)
		}
	}(time.Duration(int64(time.Minute) * int64(config.Get().ServiceCheckMinutes)))

	func() {
		for {
			select {
			case <-ctx.Done():
				stop()

				err = endBytesMap.WriteToFile(saveMapFile)
				if err != nil {
					log.Printf("Save endBytesMap to file %s, err:%s", saveMapFile, err.Error())
				}

				os.Exit(2)
			}
		}
	}()
}
