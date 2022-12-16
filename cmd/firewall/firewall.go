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
			log.Println("Blocked in ", time.Since(start).Microseconds(), ctx.Err())
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

/*
func TmpTest() {
	for i := 0; i < 20000; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		iptablesBlock(context.Background(), fmt.Sprintf("%d.%d.%d.%d",
			1+rand.Intn(255-1), 1+rand.Intn(255-1), 1+rand.Intn(255-1), 1+rand.Intn(255-1)))
	}
}
*/
