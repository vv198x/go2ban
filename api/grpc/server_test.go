package grpc

import (
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	tests := []struct {
		name        string
		port        string
		runAsDaemon bool
	}{
		{"Start and Run as Daemon", "2048/tcp", true},
		{"Start and Run as Daemon", "0/tcp", true},
		{"Start and Run as Daemon", "2048/tcp", false},
		{"Start and Not run as Daemon", "abc", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s := grpc.NewServer()
			defer s.Stop()

			RegisterIP2BanServer(s, &Server{})

			Start(tt.runAsDaemon)

			time.Sleep(time.Millisecond * 100)
		})
	}
}
