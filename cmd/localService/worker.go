package localService

import (
	"context"
	"go2ban/pkg/config"
	"go2ban/pkg/syncMap"
	"os"
	"os/signal"
	"syscall"
)

func WorkerStart(services []config.Service) {
	if config.Get().Flags.RunAsDaemon == false {
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	// if map not old read map
	//TODO add check
	countFailsMap := syncMap.NewCountersMap()
	endBytesMap := syncMap.NewStorageMap()
	for _, service := range services {
		if service.On {
			go checkLogAndBlock(ctx, service, countFailsMap, endBytesMap)
		}
	}

	func() {
		for {
			select {
			case <-ctx.Done():
				stop()
				// save map
				os.Exit(2)
			}
		}
	}()
}
