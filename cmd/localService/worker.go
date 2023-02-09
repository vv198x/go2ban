package localService

import (
	"context"
	"github.com/vv198x/go2ban/pkg/config"
	"github.com/vv198x/go2ban/pkg/docker"
	"github.com/vv198x/go2ban/pkg/storage"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

const nameMapFile = "endBytesMap"

func WorkerStart(services []config.Service, pprofEnd interface{ Stop() }) {
	//todo проверить сервисы, если все выключены не или нет выйти
	if !config.Get().Flags.RunAsDaemon {
		return
	}
	// Context for get exit signal, save map to file
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	countFailsMap := storage.NewCountersMap()
	endBytesMap := storage.NewStorageMap()

	saveMapFile := filepath.Join(config.Get().LogDir, nameMapFile)
	if err := endBytesMap.ReadFromFile(saveMapFile); err != nil {
		log.Println("Can't read from file map: ", err)
	}

	go func(sleepMinutes time.Duration) {
		for {
			for _, service := range services {
				if service.On {
					// If we get the work log of all working containers in the docker
					if service.LogFile == "docker" {
						if dockerSysLogs, err := docker.GetListsSyslogFiles(); err == nil {
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

				if err := endBytesMap.WriteToFile(saveMapFile); err != nil {
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
