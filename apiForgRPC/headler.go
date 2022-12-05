package apiForgRPC

import (
	"context"
	"go2ban/cmd/firewall"
	"go2ban/pkg/validator"
)

type Server struct {
	UnimplementedIP2BanServer
}

func (s *Server) IP(ctx context.Context, in *IPStringRequest) (*OKReply, error) {
	ip, err := validator.CheckIp(in.Ip)
	if err == nil {
		go firewall.BlockIP(ip)
		return &OKReply{Ok: true}, nil
	}
	ctx.Done()
	return &OKReply{Ok: false}, err
}
