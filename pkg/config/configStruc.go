package config

type Cfg struct {
	Flags      flags
	Firewall   string
	LogDir     string
	GrpcPort   string
	BlockedIps int
	Services   []Service `json:"Service"`
	WhiteList  []string
}

type Service struct {
	Name    string
	On      bool
	LogFile string
	Regxp   string
}

const (
	defaultCfgFile    = "/etc/go2ban/go2ban.conf"
	defaultRunDaemon  = false
	defaultLogDir     = "/var/log/go2ban"
	defaultBlockedIps = 1000
)

var exportCfg = Cfg{}

func Get() *Cfg {
	return &exportCfg
}
