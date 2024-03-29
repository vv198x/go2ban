package localService

import (
	"context"
	"github.com/vv198x/go2ban/config"
	"github.com/vv198x/go2ban/pkg/docker"
	"github.com/vv198x/go2ban/storage"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

const nameMapFile = "endBytesMap"

type serviceWork struct {
	Name        string
	FindSt      [][]byte
	SysLogFiles []string
}

func WorkerStart(mockCtx context.Context, runAsDaemon bool, services []config.Service, pprofEnd interface{ Stop() }) {
	if !runAsDaemon {
		return
	}

	// If docker add up all search strings in array
	dockerSts := make([][]byte, 0)
	servicesWork := make([]serviceWork, 0)
	for _, s := range services {
		if !s.On {
			continue
		}
		if s.LogFile != config.IsDocker {
			servicesWork = append(servicesWork, serviceWork{
				Name:        s.Name,
				FindSt:      [][]byte{[]byte(s.Regxp)},
				SysLogFiles: []string{s.LogFile},
			})
		} else { // Docker
			dockerSts = append(dockerSts, []byte(s.Regxp))
		}
	}

	// Service for docker
	if len(dockerSts) > 0 {
		if dockerSysLogs, err := docker.GetListsSyslogFiles(); err == nil {

			servicesWork = append(servicesWork, serviceWork{
				Name:        config.IsDocker,
				FindSt:      dockerSts,
				SysLogFiles: dockerSysLogs,
			})
		}
	}

	// Context for get exit signal, save map to file
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Change to a context with a timeout for the test
	if mockCtx != context.TODO() {
		ctx = mockCtx
	}

	// In-memory cache
	endBytesMap := storage.NewSyncMap()

	saveMapFile := filepath.Join(config.Get().LogDir, nameMapFile)
	if err := endBytesMap.ReadFromFile(saveMapFile); err != nil {
		log.Println("Can't read from file map: ", err)
	}

	go func(sleepMinutes time.Duration) {
		for {
			countFailsMap := storage.NewSyncMap()
			for _, sw := range servicesWork {
				for _, f := range sw.SysLogFiles {
					go sw.checkLogAndBlock(ctx, f, countFailsMap, endBytesMap)
				}
			}
			time.Sleep(sleepMinutes)
		}
	}(time.Duration(int64(time.Minute) * int64(config.Get().ServiceCheckMinutes)))

	<-ctx.Done()
	{
		if err := endBytesMap.WriteToFile(saveMapFile); err != nil {
			log.Printf("Save endBytesMap to file %s, err:%s", saveMapFile, err.Error())
		}

		if pprofEnd != nil {
			pprofEnd.Stop()
		}

		if mockCtx == context.TODO() {
			os.Exit(2)
		}
	}
}
