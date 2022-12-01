package apiForgRPC

import (
	context "context"
)

type Server struct {
	UnimplementedIP2BanServer
}

func (s *Server) IP(ctx context.Context, in *IPStringRequest) (*OKReply, error) {
	return &OKReply{Ok: true}, nil
}
