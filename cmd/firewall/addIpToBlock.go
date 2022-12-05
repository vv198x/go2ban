package firewall

import "go2ban/pkg/config"

func BlockIP(ip string) {
	switch config.Get().Firewall {
	case "firewalld":
		firewallBlock(ip)
	}
}

func UnlockAll() (blockedIp int, err error) {
	switch config.Get().Firewall {
	case "firewalld":
		return firewalldUnlockAll()
	}
	return
}
