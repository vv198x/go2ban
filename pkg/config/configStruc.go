package config

type Cfg struct {
	Flags     flags
	LogDir    string
	Firewall  string
	Services  []Service `json:"Service"`
	whiteList []string
}

type Service struct {
	Name    string
	On      bool
	LogFile string
	Regxp   string
}

const (
	defaultCfgFile   = "/etc/go2ban/go2ban.conf"
	defaultRunDaemon = false
)

// default
var exportCfg = Cfg{
	LogDir: "/var/log/go2ban",
}

func Get() *Cfg {
	return &exportCfg
}
