package config

type cfg struct {
	Flags     flags
	LogDir    string
	Firewall  string
	whiteList []string
}

const (
	defaultCfgFile   = "/etc/go2ban/go2ban.conf"
	defaultRunDaemon = false
)

// default
var exportCfg = cfg{
	LogDir: "/var/log/go2ban",
}

func Get() *cfg {
	return &exportCfg
}
