package addFirewall

import "go2ban/pkg/config"

func BlockIP(ip string) {
	switch config.Get().Firewall {
	case "firewalld":
		firewalldBlock(ip)
	}
}
