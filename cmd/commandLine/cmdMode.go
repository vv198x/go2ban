package commandLine

import (
	"context"
	"fmt"
	"github.com/vv198x/go2ban/cmd/firewall"
	"github.com/vv198x/go2ban/config"
)

func Run() {
	switch {
	case config.Get().Flags.UnlockAll:
		blockedIp, err := firewall.Do().UnlockAll(context.Background())
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
