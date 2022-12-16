package config

import (
	"encoding/json"
	"go2ban/pkg/osUtil"
	"log"
	"strconv"
	. "strings"
	"syscall"
)

func Load() {
	if syscall.Getegid() != 0 {
		log.Fatalln("Only the root user is allowed to run")
	}

	readFlags()

	if !osUtil.CheckFile(exportCfg.Flags.ConfigFile) {
		log.Fatalln("Config file not found")
	}
	cfgSt, err := osUtil.ReadStsFile(exportCfg.Flags.ConfigFile)
	if err != nil || len(cfgSt) == 0 {
		log.Fatalln("Err read config file ")
	}

	exportCfg.LogDir = defaultLogDir
	exportCfg.BlockedIps = defaultBlockedIps
	exportCfg.FakeSocksFails = defaultFakeSocksFails
	exportCfg.ServiceCheckMinutes = defaultServiceCheck
	exportCfg.ServiceFails = defaultServiceFails

	jsonData := []byte{}

	for i, line := range cfgSt {
		splitSt := Split(line, "=")
		if line[0] != byte('#') && len(splitSt) > 0 {

			switch splitSt[0] {
			case "grpc_port":
				exportCfg.GrpcPort = Fields(splitSt[1])[0]
			case "rest_port":
				exportCfg.RestPort = Fields(splitSt[1])[0]

			case "blocked_ips":
				toInt, err := strconv.Atoi(Fields(splitSt[1])[0])
				if err == nil {
					exportCfg.BlockedIps = toInt
				}

			case "log_dir":
				exportCfg.LogDir = Fields(splitSt[1])[0]

			case "firewall":
				if Contains(splitSt[1], "auto") {
					firewallName := whatFirewall()
					cfgSt[i] = Join([]string{splitSt[0], firewallName}, "=")
					err = osUtil.WriteStrsFile(cfgSt, exportCfg.Flags.ConfigFile)
					if err != nil {
						log.Println("Cant overwrite config file", err)
					}
					exportCfg.Firewall = firewallName
				} else {
					exportCfg.Firewall = Fields(splitSt[1])[0]
				}

			case "white_list":
				exportCfg.WhiteList = Fields(splitSt[1])

			case "fake_socks_ports":
				bufPorts := Fields(splitSt[1])
				for _, port := range bufPorts {
					portInt, err := strconv.Atoi(port)
					if err == nil {
						exportCfg.FakeSocksPorts = append(exportCfg.FakeSocksPorts, portInt)
					}
				}
			case "fake_socks_fails":
				fails, err := strconv.Atoi(splitSt[1])
				if err == nil {
					exportCfg.FakeSocksFails = fails
				}

			case "local_service_check_minutes":
				minutes, err := strconv.Atoi(splitSt[1])
				if err == nil {
					exportCfg.FakeSocksFails = minutes
				}
			case "local_service_fails":
				fails, err := strconv.Atoi(splitSt[1])
				if err == nil {
					exportCfg.FakeSocksFails = fails
				}
			}
		}

		if line == "{" {
			for _, jsonSt := range cfgSt[i:] {
				jsonData = append(jsonData, jsonSt[:]...)
			}

			break
		}
	}
	if len(jsonData) > 0 {
		err = json.Unmarshal(jsonData, &exportCfg)
		if err != nil {
			log.Println("Wrong json format in config file", err)
		}
	}
}

func whatFirewall() (firewallType string) {
	if osUtil.CheckFile("/usr/sbin/iptables") {
		return "iptables"
	} else {
		log.Fatalln("iptables not found")
	}
	/*systemdEnableServiceDir := "/etc/systemd/system/multi-user.target.wants/"
	firewalls := []string{
		"firewalld", //"ufw",//"shorewall",
		"iptables",
	}
	for _, firewall := range firewalls {
		serviceFile := filepath.Join(systemdEnableServiceDir, firewall+".service")
		if osUtil.CheckFile(serviceFile) {
			return firewall
		}
	}
	log.Fatalln("Firewall not found")*/
	return
}
