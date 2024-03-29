package firewall

import (
	"context"
	"github.com/vv198x/go2ban/config"
	"log"
	"os/exec"
)

type Firewall interface {
	Block(ctx context.Context, ip string)
	Worker()
	UnlockAll(ctx context.Context) (ips int, err error)
	countBlocked() (ips int)
	GetBlocked() map[string]struct{}
}

var ExportFirewall Firewall

func Do() Firewall {
	return ExportFirewall
}

func Initialization(runAsDaemon bool) {
	switch config.Get().Firewall {
	case config.IsIptables:
		ExportFirewall = &iptables{}
	case config.IsMock:
		ExportFirewall = &Mock{}
	default:
		log.Fatalln("Bad firewall")
	}

	if !runAsDaemon {
		return
	}

	ExportFirewall.Worker()
}

func runCMD(firewallCMD string) error {
	return exec.Command("sh", "-c", firewallCMD).Run()
}

func runOutputCMD(firewallCMD string) ([]byte, error) {
	return exec.Command("sh", "-c", firewallCMD).Output()
}
