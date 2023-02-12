package grpc

import (
	"context"
	"github.com/vv198x/go2ban/cmd/firewall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIP(t *testing.T) {
	firewall.ExportFirewall = &firewall.Mock{}
	server := &Server{}

	t.Run("Test case 1: valid IP address", func(t *testing.T) {
		ip := "192.168.1.1"
		in := &IPStringRequest{Ip: ip}
		out, err := server.IP(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, true, out.Ok)
	})

	t.Run("Test case 1: Not valid IP address", func(t *testing.T) {
		ip := "invalid ip"
		in := &IPStringRequest{Ip: ip}
		out, err := server.IP(context.Background(), in)
		assert.NotNil(t, err)
		assert.Equal(t, false, out.Ok)
	})
}
