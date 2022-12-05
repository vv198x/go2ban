package firewall

import (
	"go2ban/pkg/config"
)

func BlockIP(ip string) {
	switch config.Get().Firewall {
	case "firewalld":
		firewallBlock(ip)
	}
}

func UnlockAll() (blockedIp int, err error) {
	switch config.Get().Firewall {
	case "firewalld":
		return firewalldUnlockAll()
	}
	return
}

/*
func TmpTest() {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		firewallBlock(fmt.Sprintf("%d.%d.%d.%d",
			1+rand.Intn(255-1), 1+rand.Intn(255-1), 1+rand.Intn(255-1), 1+rand.Intn(255-1)))

	}
}
*/
