package commandLine

import (
	"bytes"
	"github.com/vv198x/go2ban/pkg/config"
	"io"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	// Store the original value of stdout
	rescueStdout := os.Stdout

	// Create a pipe for capturing the output of the function
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Set the configuration flag to true (-clear)
	config.Get().Flags.UnlockAll = true

	// Call the Run function
	Run()

	// Close the write end of the pipe
	w.Close()

	// Copy the captured output to a buffer
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Restore the original value of stdout
	os.Stdout = rescueStdout

	// Verify the output
	if !strings.Contains(buf.String(), "IPs unlocked") {
		t.Errorf("Expected %s, got %s", "IPs unlocked", buf.String())
	}
	r.Close()
}
