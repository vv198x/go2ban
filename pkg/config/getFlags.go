package config

import "flag"

type flags struct {
	RunAsDaemon bool
	UnlockAll   bool
	ConfigFile  string
}

func ReadFlags() {
	cfgFile := flag.String("cfgFile", defaultCfgFile, "Path to file go2ban.conf")
	daemon := flag.Bool("d", defaultRunDaemon, "Run as daemon")
	clear := flag.Bool("clear", false, "Unlock all")
	flag.Parse()
	exportCfg.Flags.RunAsDaemon = *daemon
	exportCfg.Flags.ConfigFile = *cfgFile
	exportCfg.Flags.UnlockAll = *clear
}
