package osUtil

import (
	"net"
)

var l []string

func GetLocalIPs() []string {
	if len(l) > 0 {
		return l
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return l
	}

	for _, address := range addrs {
		ipIntf := address.(*net.IPNet).IP.To4()
		if ipIntf != nil {
			l = append(l, ipIntf.String())
		}
	}
	return l
}
