package osUtil

import (
	"net"
	"testing"
)

func TestGetLocalIPs(t *testing.T) {
	l = nil

	ips := GetLocalIPs()

	if len(ips) == 0 {
		t.Fatalf("No IP addresses returned")
	}

	var localhost bool
	for _, ip := range ips {
		if net.ParseIP(ip) == nil {
			t.Fatalf("Invalid IP address returned: %s", ip)
		}
		if ip == "127.0.0.1" {
			localhost = !localhost
		}
	}
	if !localhost {
		t.Fatalf("Local IP address not found")
	}
}
