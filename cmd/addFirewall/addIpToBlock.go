package addFirewall

import (
	"go2ban/pkg/osUtil"
	"log"
	"path/filepath"
)

func BlockIP(ip string) {
	switch whatFirewall() {
	case "firewalld":
		firewalldBlock(ip)
	}
}

func whatFirewall() (firewallType string) {
	systemdEnableServiseDir := "/etc/systemd/system/multi-user.target.wants/"
	firewalls := []string{
		"firewalld",
		//"ufw",
		//"shorewall",
		"iptables",
	}
	for _, firewall := range firewalls {
		serviceFile := filepath.Join(systemdEnableServiseDir, firewall+".service")
		if osUtil.CheckFile(serviceFile) {
			return firewall
		}
	}
	log.Fatalln("Firewall not found")
	return
}
