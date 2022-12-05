package config

import . "flag"

type flags struct {
	RunAsDaemon bool
	UnlockAll   bool
	ConfigFile  string
}

func init() {
	cfgFile := String("cfgFile", defaultCfgFile, "Path to file go2ban.conf")
	daemon := Bool("d", defaultRunDaemon, "Run as daemon")
	clear := Bool("clear", false, "Remove all rules")
	Parse()
	exportCfg.Flags.ConfigFile = *cfgFile
	exportCfg.Flags.RunAsDaemon = *daemon
	exportCfg.Flags.UnlockAll = *clear
}
