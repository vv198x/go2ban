package firewall

import (
	"context"
	"go2ban/pkg/config"
	"log"
	"time"
)

func BlockIP(ctx context.Context, ip string) {
	switch config.Get().Firewall {
	case "iptables":
		iptablesBlock(ctx, ip)
	}
	go func() {
		select {
		case <-time.After(10 * time.Microsecond):
			log.Println("* Longer than average *")
		}
	}()
}

func UnlockAll() (blockedIp int, err error) {
	switch config.Get().Firewall {
	case "iptables":
		return firewalldUnlockAll()
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
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		iptablesBlock(fmt.Sprintf("%d.%d.%d.%d",
			1+rand.Intn(255-1), 1+rand.Intn(255-1), 1+rand.Intn(255-1), 1+rand.Intn(255-1)))

	}
}
*/
