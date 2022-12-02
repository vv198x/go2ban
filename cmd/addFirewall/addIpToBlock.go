package addFirewall

import (
	"fmt"
	"microservice2ban/pkg/osUtil"
	"path/filepath"
)

func BlockIP(ip string) {
	fmt.Println(whatFirewall())
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
}
