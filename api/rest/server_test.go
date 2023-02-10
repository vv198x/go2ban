package rest

import (
	"github.com/vv198x/go2ban/config"
	"testing"
)

func TestStart(t *testing.T) {
	tests := []struct {
		name           string
		runAsDaemon    bool
		configRestPort string
		wantErr        bool
	}{
		{"Test case 1: run as daemon, valid rest port", true, "8080", false},
		{"Test case 2: run as daemon, invalid rest port", true, "abc", true},
		{"Test case 3: not run as daemon", false, "8080", false},
		{"Test case 4: run as daemon, empty rest port", true, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.Get().RestPort = tt.configRestPort
			defer func() {
				if r := recover(); r != nil {
					if tt.wantErr {
						return
					}
					t.Errorf("Unexpected panic: %v", r)
				}
			}()
			Start(tt.runAsDaemon)
		})
	}
}
