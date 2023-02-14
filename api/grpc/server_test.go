package grpc

import (
	"bytes"
	"github.com/vv198x/go2ban/config"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
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

func TestStart2(t *testing.T) {
	// Test case 1: runAsDaemon=false
	Start(false) // Expect no action

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	// Test case 2: valid grpc port configuration and runAsDaemon=true
	config.Get().GrpcPort = "8080/tcp"
	runAsDaemon := true
	Start(runAsDaemon)
	time.Sleep(100 * time.Millisecond) // wait for goroutine to start
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		t.Errorf("Failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	// Test case 3: invalid grpc port configuration and runAsDaemon=true
	config.Get().GrpcPort = "8080" // invalid port format
	Start(runAsDaemon)
	time.Sleep(100 * time.Millisecond) // wait for goroutine to start
	// Expect error message logged to console

	Start(runAsDaemon)

	if strings.Contains(buf.String(), "gRPC error") {
		t.Errorf("Expected error message for invalid port configuration")
	}

}
