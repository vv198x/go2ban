package firewall

import (
	"context"
	"go2ban/pkg/config"
	"log"
	"time"
)

const sleepHour = 12

func BlockIP(ctx context.Context, ip string) {
	switch config.Get().Firewall {
	case "iptables":
		iptablesBlock(ctx, ip)
	}
	start := time.Now()
	go func() {
		select {
		case <-ctx.Done():
			log.Println("Blocked in microseconds :", time.Since(start).Microseconds())
		case <-time.After(50 * time.Microsecond):
			log.Println("* Runs longer than usual *")
		}
	}()
}

func UnlockAll(ctx context.Context) (blockedIp int, err error) {
	switch config.Get().Firewall {
	case "iptables":
		return iptablesUnlockAll(ctx)
	}
	return
}

func WorkerStart(RunAsDaemon bool) {
	if RunAsDaemon {
		switch config.Get().Firewall {
		case "iptables":
			workerIptables()
		}
	}
}
