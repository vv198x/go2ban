package fakeSocks

import (
	"context"
	"fmt"
	"go2ban/cmd/firewall"
	"go2ban/pkg/config"
	"go2ban/pkg/validator"
	"log"
	"net"
	"strconv"
)

func Listen(ports []int) {
	for _, port := range ports {
		p := strconv.Itoa(port)
		p = ":" + p
		go func() {
			listener, err := net.Listen("tcp", p)
			if err != nil {
				log.Println("Fake socks: don't listen", err)
				return
			}
			log.Println("Fake socks open port", p)
			defer listener.Close()

			cuntMap := newCounters()

			for {
				conn, err := listener.Accept()
				defer conn.Close()
				if err != nil {
					fmt.Println("Fake socks listener", p, err)
					continue
				}

				badIp, err := validator.CheckIp(conn.RemoteAddr().String())
				if err != nil {
					log.Println("Fake socks error addr", p, err)
					continue
				}

				cuntMap.Inc(badIp)
				if cuntMap.Load(badIp) >= config.Get().FakeSocksFails {
					ctx := context.Background()
					go firewall.BlockIP(ctx, badIp)
					log.Println("Fake socks ip bloked:", badIp)
				}
			}
		}()
	}
}
