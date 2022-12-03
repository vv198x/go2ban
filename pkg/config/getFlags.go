package config

import (
	"flag"
)

func init() {
	exportCfg.ConfigFile = *flag.String("cfgFile", defaultCfgFile, "Path to file go2ban.conf")
	exportCfg.RunAsDaemon = *flag.Bool("d", defaultRunDaemon, "Run as daemon")
	flag.Parse()
}
