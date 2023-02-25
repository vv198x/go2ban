package config

type Cfg struct {
	Flags               flags
	Firewall            string
	LogDir              string
	GrpcPort            string
	RestPort            string
	BlockedIps          int
	AbuseipdbApiKey     string
	AbuseipdbIPs        int
	WhiteList           []string
	TrapPorts           []int
	TrapFails           int
	ServiceCheckMinutes int
	ServiceFails        int
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
	defaultAbuseipdbIPs   = 2000
	IsDocker              = "docker"
	IsIptables            = "iptables"
	IsMock                = "mock"
	WorkerSleepHour       = 6
)

var exportCfg = Cfg{}

func Get() *Cfg {
	return &exportCfg
}
