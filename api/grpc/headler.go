package grpc

import (
	"context"
	"github.com/vv198x/go2ban/cmd/firewall"
	"github.com/vv198x/go2ban/pkg/validator"
	"log"
)

type Server struct { //nolint
	UnimplementedIP2BanServer
}

// nolint
func (s *Server) IP(ctx context.Context, in *IPStringRequest) (*OKReply, error) {
	ip, err := validator.CheckIp(in.Ip)

	if err == nil {
		go firewall.BlockIP(ctx, ip)
		log.Println("gRPC ip blocked:", ip)
		return &OKReply{Ok: true}, nil

	} else {
		log.Println("gRPC validator:", err)
	}
	ctx.Done()
	return &OKReply{Ok: false}, err
}
