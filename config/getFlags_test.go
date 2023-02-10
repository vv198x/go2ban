package config

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestReadFlags(t *testing.T) {
	tests := []struct {
		name    string
		cfgFile string
		daemon  bool
		clear   bool
	}{
		{"default values", defaultCfgFile, defaultRunDaemon, false},
		{"custom values", "/tmp/test.conf", false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flags before each test
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

			// Set custom arguments
			os.Args = []string{"cmd", fmt.Sprintf("--cfgFile=%s", tt.cfgFile), fmt.Sprintf("-d=%t", tt.daemon), fmt.Sprintf("--clear=%t", tt.clear)}
			ReadFlags()

		})
	}
}
