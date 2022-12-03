package config

type cfg struct {
	RunAsDaemon bool
	ConfigFile  string
	LogDir      string
	firewall    string
	whiteList   []string
}

const (
	defaultCfgFile   = "/etc/go2ban/go2ban.conf"
	defaultRunDaemon = false
)

var exportCfg = cfg{
	LogDir: "/var/log/go2ban",
}

func Get() *cfg {
	return &exportCfg
}
