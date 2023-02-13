package firewall

import (
	"fmt"
	"github.com/vv198x/go2ban/config"
	"reflect"
	"testing"
)

func TestRunOutputCMD(t *testing.T) {
	firewallCMD := "echo hello"
	output, err := runOutputCMD(firewallCMD)
	if err != nil {
		t.Errorf("Expected error to be nil but got %v, output %v", err, output)
	}
	if string(output) != "hello\n" {
		t.Errorf("Expected output to be 'hello\n' but got %s", output)
	}

	firewallCMD = "invalid_command"
	output, err = runOutputCMD(firewallCMD)
	if err == nil {
		t.Errorf("Expected error but got nil, output: %v", output)
	} else {
		fmt.Printf("firewallCMD: %s\n", firewallCMD)
		fmt.Printf("error: %v\n", err)
	}
}
func TestRunCMD(t *testing.T) {
	firewallCMD := "echo hello"
	err := runCMD(firewallCMD)
	if err != nil {
		t.Errorf("Expected error to be nil but got %v", err)
	}

	firewallCMD = "invalid_command"
	err = runCMD(firewallCMD)
	if err == nil {
		t.Errorf("Expected error but got nil")
	} else {
		fmt.Printf("firewallCMD: %s\n", firewallCMD)
		fmt.Printf("error: %v\n", err)
	}
}

func TestInitialization(t *testing.T) {
	tests := []struct {
		name           string
		configFirewall string
		runAsDaemon    bool
		wantFirewall   string
	}{
		{"Test mock", config.IsMock, false, "*firewall.Mock"},
		{"Test iptables", config.IsIptables, true, "*firewall.iptables"},
	}
	backFW := config.Get().Firewall
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.Get().Firewall = tt.configFirewall
			Initialization(tt.runAsDaemon)
			if fmt.Sprintf("%T", ExportFirewall) != tt.wantFirewall {
				t.Errorf("Initialization() = %v, want %v", reflect.TypeOf(ExportFirewall), tt.wantFirewall)
			}
		})
	}
	config.Get().Firewall = backFW
}
