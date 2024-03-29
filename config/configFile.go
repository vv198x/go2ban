package config

import (
	"encoding/json"
	"github.com/vv198x/go2ban/pkg/osUtil"
	"log"
	"strconv"
	"strings"
	"syscall"
)

func Load() {
	if syscall.Getegid() != 0 {
		log.Fatalln("Only the root user is allowed to run")
	}

	if !osUtil.CheckFile(exportCfg.Flags.ConfigFile) {
		log.Fatalf("Config file not found, file: %v", exportCfg.Flags.ConfigFile)
	}

	cfgSt, err := osUtil.ReadStsFile(exportCfg.Flags.ConfigFile)
	if err != nil || len(cfgSt) == 0 {
		log.Fatalln("Err read config file ")
	}

	exportCfg.LogDir = defaultLogDir
	exportCfg.BlockedIps = defaultBlockedIps
	exportCfg.TrapFails = defaultFakeSocksFails
	exportCfg.ServiceCheckMinutes = defaultServiceCheck
	exportCfg.ServiceFails = defaultServiceFails
	exportCfg.AbuseipdbIPs = defaultAbuseipdbIPs

	jsonData := make([]byte, 0)

	for i, line := range cfgSt {
		if len(line) == 0 {
			continue
		}
		splitSt := strings.Split(line, "=")
		if line[0] != byte('#') && len(splitSt) > 0 {
			// Start JSON data
			if line == "{" {
				for _, jsonSt := range cfgSt[i:] {
					jsonData = append(jsonData, jsonSt[:]...)
				}
				// JSON data at the end of the file, break the loop
				break
			}

			params := strings.Fields(splitSt[1])

			switch splitSt[0] {
			case "grpc_port":
				exportCfg.GrpcPort = params[0]
			case "rest_port":
				exportCfg.RestPort = params[0]

			case "blocked_ips":
				toInt, err := strconv.Atoi(params[0])
				if err == nil {
					exportCfg.BlockedIps = toInt
				}

			case "log_dir":
				exportCfg.LogDir = params[0]

			case "firewall":
				if strings.Contains(splitSt[1], "auto") {
					firewallName := whatFirewall()
					cfgSt[i] = strings.Join([]string{splitSt[0], firewallName}, "=")
					err = osUtil.WriteStrsFile(cfgSt, exportCfg.Flags.ConfigFile)
					if err != nil {
						log.Println("Cant overwrite config file", err)
					}
					exportCfg.Firewall = firewallName
				} else {
					exportCfg.Firewall = params[0]
				}

			case "white_list":
				exportCfg.WhiteList = params

			case "trap_ports":
				bufPorts := params
				for _, port := range bufPorts {
					portInt, err := strconv.Atoi(port)
					if err == nil {
						exportCfg.TrapPorts = append(exportCfg.TrapPorts, portInt)
					}
				}
			case "trap_fails":
				fails, err := strconv.Atoi(params[0])
				if err == nil {
					exportCfg.TrapFails = fails
				}

			case "local_service_check_minutes":
				minutes, err := strconv.Atoi(params[0])
				if err == nil {
					exportCfg.ServiceCheckMinutes = minutes
				}
			case "local_service_fails":
				fails, err := strconv.Atoi(params[0])
				if err == nil {
					exportCfg.ServiceFails = fails
				}

			case "abuseipdb_apikey":
				exportCfg.AbuseipdbApiKey = params[0]
			case "abuseipdb_ips":
				ips, err := strconv.Atoi(params[0])
				if err == nil {
					exportCfg.AbuseipdbIPs = ips
				}
			}
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
	return
}
