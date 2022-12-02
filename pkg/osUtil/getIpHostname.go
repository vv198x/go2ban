package osUtil

import (
	"log"
	"net"
	"os"
)

func GetLocalIPs() (end []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalln("It's impossible error")
		return
	}
	for _, address := range addrs {
		ipIntf := address.(*net.IPNet).IP.To4()
		if ipIntf != nil {
			end = append(end, ipIntf.String())
		}

	}
	return
}

func GetLocalHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Println("No get localhostname ", err)
		return ""
	}
	return hostname
}
