package config

type cfg struct {
	RunAsDaemon bool
	ConfigFile  string
	LogDir      string
	Firewall    string
	whiteList   []string
}

const (
	defaultCfgFile   = "/etc/go2ban/go2ban.conf"
	defaultRunDaemon = false
)

// default
var exportCfg = cfg{
	LogDir:   "/var/log/go2ban",
	Firewall: "firewalld",
}

func Get() *cfg {
	return &exportCfg
}
