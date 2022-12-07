package firewall

import (
	"go2ban/pkg/config"
)

func WorkerStart(RunAsDaemon bool) {
	if RunAsDaemon {
		switch config.Get().Firewall {
		case "iptables":
			workerIptables()
		}
	}
}
