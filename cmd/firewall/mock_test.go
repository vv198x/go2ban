package firewall

import (
	"bytes"
	"context"
	"log"
	"os"
	"strings"
	"testing"
)

func TestMock_Block(t *testing.T) {
	type args struct {
		ctx context.Context
		ip  string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test case 1", args{context.Background(), "192.168.0.1"}},
		{"Test case 2", args{context.Background(), "192.168.0.2"}},
	}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fw := &Mock{}
			fw.Block(tt.args.ctx, tt.args.ip)

			// check if the log message was printed as expected
			if !strings.Contains(buf.String(), "Mock firewall blocked ip: "+tt.args.ip) {
				t.Errorf("Expected log message not found")
			}
		})
	}
}

func TestMock_UnlockAll(t *testing.T) {
	mock := &Mock{}
	ctx := context.Background()

	ips, err := mock.UnlockAll(ctx)

	if err != nil {
		t.Fatalf("Expected error to be nil, but got: %v", err)
	}
	if ips != 0 {
		t.Fatalf("Expected number of unlocked IPs to be 0, but got: %d", ips)
	}
}

func TestMock_countBlocked(t *testing.T) {
	mock := &Mock{}

	ips := mock.countBlocked()

	if ips != 0 {
		t.Fatalf("Expected number of blocked IPs to be 0, but got: %d", ips)
	}
}
