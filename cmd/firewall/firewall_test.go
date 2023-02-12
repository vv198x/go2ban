package firewall

import (
	"fmt"
	"testing"
)

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
