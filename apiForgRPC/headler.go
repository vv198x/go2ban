package apiForgRPC

import (
	"context"
	"microservice2ban/pkg/validator"
)

type Server struct {
	UnimplementedIP2BanServer
}

func (s *Server) IP(ctx context.Context, in *IPStringRequest) (*OKReply, error) {
	err := validator.CheckIp(in.Ip)
	if err == nil {
		return &OKReply{Ok: true}, nil
	}
	return &OKReply{Ok: false}, err
}
