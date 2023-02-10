package config

import (
	"flag"
)

type flags struct {
	RunAsDaemon bool
	UnlockAll   bool
	ConfigFile  string
}

func ReadFlags() {
	cfgFile := flag.String("cfgFile", defaultCfgFile, "Path to file go2ban.conf")
	daemon := flag.Bool("d", defaultRunDaemon, "Run as daemon")
	clear := flag.Bool("clear", false, "Unlock all")

	_ = flag.Bool("test.v", false, "Test flag")
	_ = flag.Bool("test.paniconexit0", false, "Test flag")
	_ = flag.Bool("test.run", false, "Test flag")

	flag.Parse()
	exportCfg.Flags.RunAsDaemon = *daemon
	exportCfg.Flags.ConfigFile = *cfgFile
	exportCfg.Flags.UnlockAll = *clear
}
