package localService

import (
	"context"
	"go2ban/pkg/config"
	"go2ban/pkg/docker"
	"go2ban/pkg/storage"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

const nameMapFile = "endBytesMap"

func WorkerStart(services []config.Service, pprofEnd interface{ Stop() }) {
	if config.Get().Flags.RunAsDaemon == false {
		return
	}
	//Контекст для, получение сигнала выхода, сохранить мапу в файл
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	countFailsMap := storage.NewCountersMap()
	endBytesMap := storage.NewStorageMap()

	saveMapFile := filepath.Join(config.Get().LogDir, nameMapFile)
	err := endBytesMap.ReadFromFile(saveMapFile)
	if err != nil {
		log.Println("Can't read from file map: ", err)
	}

	go func(sleepMinutes time.Duration) {
		for {
			for _, service := range services {
				if service.On {
					//Если в лог в докере получаем суслог всех рабочих контейнеров
					if service.LogFile == "docker" {
						dockerSysLogs, errD := docker.GetListsSyslogFiles()
						if errD == nil {
							for _, f := range dockerSysLogs {
								service.LogFile = f

								go checkLogAndBlock(ctx, service, countFailsMap, endBytesMap)
							}
						}
					} else {

						go checkLogAndBlock(ctx, service, countFailsMap, endBytesMap)
					}
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
				if pprofEnd != nil {
					pprofEnd.Stop()
				}
				os.Exit(2)
			default:
				time.Sleep(time.Second)
			}
		}
	}()
}
