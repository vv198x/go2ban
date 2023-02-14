package rest

import (
	"github.com/vv198x/go2ban/config"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	tests := []struct {
		name           string
		runAsDaemon    bool
		configRestPort string
	}{
		{"Test case 1: run as daemon, valid rest port", true, "8080"},
		{"Test case 2: run as daemon, invalid rest port", true, "abc"},
		{"Test case 3: not run as daemon", false, "8080"},
		{"Test case 4: run as daemon, empty rest port", true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.Get().RestPort = tt.configRestPort
			Start(tt.runAsDaemon)
			time.Sleep(time.Millisecond * 50)
		})
	}
}
