package localService

import (
	"context"
	"go2ban/pkg/config"
	"os"
	"os/signal"
	"syscall"
)

func WorkerStart(service []config.Service) {
	if config.Get().Flags.RunAsDaemon == false {
		return
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	func() {
		for {
			select {
			case <-ctx.Done():
				stop()
				os.Exit(2)
			}
		}
	}()
}
