package fakeSocks

import (
	"context"
	"fmt"
	"go2ban/cmd/firewall"
	"go2ban/pkg/config"
	"go2ban/pkg/syncMap"
	"go2ban/pkg/validator"
	"log"
	"net"
	"strconv"
)

func Listen(ports []int) {
	if config.Get().Flags.RunAsDaemon == false {
		return
	}

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

			countMap := syncMap.NewCountersMap()

			for {
				conn, err := listener.Accept()
				defer conn.Close()
				if err != nil {
					fmt.Println("Fake socks listener", p, err)
					continue
				}

				ip, err := validator.CheckIp(conn.RemoteAddr().String())
				if err != nil {
					log.Println("Fake socks error addr", p, err)
					continue
				}

				func(counterMap syncMap.SyncMap) {
					counterMap.Increment(ip)
					if int(counterMap.Load(ip))+1 == config.Get().FakeSocksFails {

						go firewall.BlockIP(context.Background(), ip)

						log.Println("Fake socks ip bloked:", ip)
					}
				}(countMap)
			}
		}()
	}
}
