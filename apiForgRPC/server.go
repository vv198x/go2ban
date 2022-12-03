package apiForgRPC

import (
	"google.golang.org/grpc"
	"net"
)

func Start(protocol, address string, runAsDaemon bool) {
	if runAsDaemon {
		lis, err := net.Listen(protocol, address)
		s := grpc.NewServer()
		RegisterIP2BanServer(s, &Server{})
		err = s.Serve(lis)
		if err != nil {
			panic(err)
		}
	}
}
