package fakeSocks

import (
	"context"
	"go2ban/cmd/firewall"
	"go2ban/pkg/config"
	"go2ban/pkg/storage"
	"go2ban/pkg/validator"
	"log"
	"net"
	"strconv"
	"time"
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

			countMap := storage.NewRedis()

			for {
				conn, err := listener.Accept()
				if err != nil {
					log.Fatalln("Fake socks listener", p, err)
					continue
				}

				ip, err := validator.CheckIp(conn.RemoteAddr().String())
				if err != nil {
					log.Println("Fake socks error addr", p, err)
					conn.Close()
					continue
				}

				func(counterMap storage.SyncMap) {
					counterMap.Increment(ip)
					count := int(counterMap.Load(ip))
					if count == config.Get().FakeSocksFails {

						go firewall.BlockIP(context.Background(), ip)

						log.Println("Fake socks ip bloked:", ip)
					}
				}(countMap)

				time.Sleep(time.Second)
				conn.Close()
			}
			countMap.Close()
		}()
	}
}
