package localService

import (
	"context"
	"go2ban/pkg/config"
	"go2ban/pkg/countSyncMap"
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
	countFailsMap := countSyncMap.NewCounters()
	for _, service := range services {
		if service.On {
			go checkLogAndBlock(ctx, service, countFailsMap, config.Get().SrviceFails)
		}
	}

	func() {
		for {
			select {
			case <-ctx.Done():
				// save map
				stop()
				os.Exit(2)
			}
		}
	}()
}
