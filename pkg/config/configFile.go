package config

import (
	"go2ban/pkg/osUtil"
	"log"
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
	for _, line := range cfgSt {
		splitSt := strings.Split(line, "=")
		if line[0] != byte('#') && len(splitSt) > 0 {
			switch splitSt[0] {
			case "log_dir":
				exportCfg.LogDir = splitSt[1]
			}
		}
	}

}
