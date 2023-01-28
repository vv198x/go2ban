package api2Ban

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func SendIp2BanGrpc(grpcPort string, ip string) bool {
	opt := grpc.WithInsecure()
	cc, err := grpc.Dial(grpcPort, opt)
	if err != nil {
		log.Println("Socket gRPC ", err)
	}
	defer cc.Close()
	client := NewIP2BanClient(cc)
	req := IPStringRequest{Ip: ip}
	response, err := client.IP(context.Background(), &req)
	if err != nil {
		log.Println("go2Ban say err ", err)
	} else {
		return response.Ok
	}
	return false
}
