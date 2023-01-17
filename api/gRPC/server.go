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
			split := strings.Split(config.Get().GrpcPort, "/") //2048/tcp
			if len(split) > 1 {
				lis, err := net.Listen(split[1], ":"+split[0])
				if err != nil {
					log.Fatalln("gRPC error ", err)
				}
				s := grpc.NewServer()
				RegisterIP2BanServer(s, &Server{})

				if err = s.Serve(lis); err != nil {
					log.Fatalln("gRPC error ", err)
				}

				log.Println("gRPC listen on port:", config.Get().GrpcPort)
			}
		}()

	}
}
