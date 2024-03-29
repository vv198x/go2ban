package commandLine

import (
	"bytes"
	"github.com/vv198x/go2ban/cmd/firewall"
	"github.com/vv198x/go2ban/config"
	"io"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	firewall.ExportFirewall = &firewall.Mock{}
	// Store the original value of stdout
	rescueStdout := os.Stdout

	// Create a pipe for capturing the output of the function
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Set the configuration flag to true (-clear)
	config.Get().Flags.UnlockAll = true

	// Call the Run function
	Run()

	// Close the writer end of the pipe
	if err := w.Close(); err != nil {
		t.Errorf("")
	}

	// Copy the captured output to a buffer
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Errorf("")
	}

	// Restore the original value of stdout
	os.Stdout = rescueStdout

	// Verify the output. Mock err
	if !strings.Contains(buf.String(), "Mock err") {
		t.Errorf("Expected %s, got %s", "IPs unlocked", buf.String())
	}
}
