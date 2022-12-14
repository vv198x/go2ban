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
	go func() {
		select {
		case <-time.After(30 * time.Microsecond):
			log.Println("* Runs longer than usual *")
		}
	}()
}

func UnlockAll() (blockedIp int, err error) {
	switch config.Get().Firewall {
	case "iptables":
		return iptablesUnlockAll()
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
