package gRPC

import (
	"go2ban/pkg/config"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

func Start(runAsDaemon bool) {
	if runAsDaemon {
		go func() {
			//TODO check port
			split := strings.Split(config.Get().GrpcPort, "/") //2048/tcp
			if len(split) > 1 {
				lis, err := net.Listen(split[1], ":"+split[0])
				s := grpc.NewServer()
				RegisterIP2BanServer(s, &Server{})

				err = s.Serve(lis)
				if err != nil {
					log.Fatalln("gRPC error ", err)
				}
				log.Println("gRPC listen on port:", config.Get().GrpcPort)
			}
		}()

	}
}
