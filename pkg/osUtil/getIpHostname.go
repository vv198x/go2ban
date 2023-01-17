package osUtil

import (
	"net"
)

func GetLocalIPs() (end []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
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
