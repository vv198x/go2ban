package config

type Cfg struct {
	Flags               flags
	Firewall            string
	LogDir              string
	GrpcPort            string
	RestPort            string
	BlockedIps          int
	WhiteList           []string
	FakeSocksPorts      []int
	FakeSocksFails      int
	ServiceCheckMinutes int
	SrviceFails         int
	Services            []Service `json:"Service"`
}

type Service struct {
	Name    string
	On      bool
	LogFile string
	Regxp   string
}

const (
	defaultCfgFile        = "/etc/go2ban/go2ban.conf"
	defaultRunDaemon      = false
	defaultLogDir         = "/var/log/go2ban"
	defaultBlockedIps     = 1000
	defaultFakeSocksFails = 2
	defaultServiceCheck   = 2
	defaultServiceFails   = 2
)

var exportCfg = Cfg{}

func Get() *Cfg {
	return &exportCfg
}
