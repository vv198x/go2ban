package main

import (
	"google.golang.org/grpc"
	"microservice2ban/apiForgRPC"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":2048")
	s := grpc.NewServer()
	apiForgRPC.RegisterIP2BanServer(s, &apiForgRPC.Server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
