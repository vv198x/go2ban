package commandLine

import (
	"fmt"
	"go2ban/cmd/firewall"
	"go2ban/pkg/config"
)

func Run() {
	switch {
	case config.Get().Flags.UnlockAll:
		blockedIp, err := firewall.UnlockAll()
		if err != nil {
			if blockedIp == 0 {
				fmt.Println(err)
			} else {
				fmt.Printf("%d IPs blocked. Unlock error: %v", blockedIp, err)
			}
		} else {
			fmt.Printf("%d IPs unlocked.", blockedIp)
		}

	}
}
