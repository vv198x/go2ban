package config

import (
	"go2ban/pkg/osUtil"
	"log"
	"path/filepath"
	"strings"
)

func Load() {
	if !osUtil.CheckFile(exportCfg.ConfigFile) {
		log.Fatalln("Config file not found")
	}
	cfgSt, err := osUtil.ReadStsFile(exportCfg.ConfigFile)
	if err != nil || len(cfgSt) == 0 {
		log.Fatalln("Err read config file", err)
	}
	for i, line := range cfgSt {
		splitSt := strings.Split(line, "=")
		if line[0] != byte('#') && len(splitSt) > 0 {
			switch splitSt[0] {
			case "log_dir":
				exportCfg.LogDir = splitSt[1]
			case "firewall":
				if strings.Contains(splitSt[1], "auto") {
					firewallName := whatFirewall()
					cfgSt[i] = strings.Join([]string{splitSt[0], firewallName}, "=")
					err = osUtil.WriteStrsFile(cfgSt, exportCfg.ConfigFile)
					if err != nil {
						log.Println("Cant write config", err)
					}
					exportCfg.Firewall = firewallName
				} else {
					exportCfg.Firewall = splitSt[1]
				}
			}
		}
	}
}

func whatFirewall() (firewallType string) {
	systemdEnableServiseDir := "/etc/systemd/system/multi-user.target.wants/"
	firewalls := []string{
		"firewalld", //"ufw",//"shorewall",
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
