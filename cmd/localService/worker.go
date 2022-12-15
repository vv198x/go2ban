package localService

import (
	"context"
	"fmt"
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
		fmt.Println(err)
	}

	go func(sleepMinutes int) {
		for {
			for _, service := range services {
				if service.On {

					go checkLogAndBlock(ctx, service, countFailsMap, endBytesMap)

				}
			}
			time.Sleep(time.Duration(int64(time.Minute) * int64(sleepMinutes)))
		}
	}(config.Get().ServiceCheckMinutes)

	func() {
		for {
			select {
			case <-ctx.Done():
				stop()
				err := endBytesMap.WriteToFile(saveMapFile)
				log.Printf("Save endBytesMap to file %s, err:%s", saveMapFile, err.Error())

				os.Exit(2)
			}
		}
	}()
}
