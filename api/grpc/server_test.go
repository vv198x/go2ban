package grpc

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	tests := []struct {
		name        string
		runAsDaemon bool
	}{
		{"Start and Run as Daemon", true},
		{"Start and Not run as Daemon", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lis, err := net.Listen("tcp", ":2048")
			if err != nil {
				log.Fatalln("gRPC error ", err)
			}

			defer lis.Close() //nolint

			s := grpc.NewServer()
			defer s.Stop()

			RegisterIP2BanServer(s, &Server{})

			go func() {
				Start(tt.runAsDaemon)
				if tt.runAsDaemon {
					if err = s.Serve(lis); err != nil {
						log.Fatalln("gRPC error ", err)
					}
				}
			}()

			time.Sleep(time.Second)
		})
	}
}
