package config

import "flag"

func init() {
	cfgFile := flag.String("cfgFile", defaultCfgFile, "Path to file go2ban.conf")
	daemon := flag.Bool("d", defaultRunDaemon, "Run as daemon")
	//TODO clear := flag.Bool("clear", false, "Remove all rules")

	flag.Parse()
	exportCfg.ConfigFile = *cfgFile
	exportCfg.RunAsDaemon = *daemon
}
