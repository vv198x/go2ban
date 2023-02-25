package fakeSocks

import (
	"bytes"
	"github.com/vv198x/go2ban/cmd/firewall"
	"github.com/vv198x/go2ban/config"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestListen(t *testing.T) {
	firewall.ExportFirewall = &firewall.Mock{}

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Listen panicked: %v", r)
		}
	}()
	// Capture the logs
	rescueStdout := os.Stdout
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stdout)

	ports := []int{1024, 1025}

	config.Get().Flags.RunAsDaemon = true
	config.Get().TrapFails = 3
	// Start the listener in a separate goroutine
	done := make(chan struct{})
	go func() {
		Listen(ports)
		close(done)
	}()

	// Wait for the listener to start
	time.Sleep(100 * time.Millisecond)

	// Connect to the listener and make sure the connection is accepted
	for _, port := range ports {
		conn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(port))
		if err != nil {
			t.Fatalf("Failed to dial listener on port %d: %v", port, err)
		}
		conn.Close()
	}
	// Wait for the listener to stop
	<-done

	// Check the logs for the expected messages
	expectedOpen := "Fake socks open port :" + strconv.Itoa(ports[0])
	if !strings.Contains(buf.String(), expectedOpen) {
		t.Errorf("Expected log to contain %q, got %q", expectedOpen, buf.String())
	}
	expectedOpen = "Fake socks open port :" + strconv.Itoa(ports[1])
	if !strings.Contains(buf.String(), expectedOpen) {
		t.Errorf("Expected log to contain %q, got %q", expectedOpen, buf.String())
	}
	//fmt.Println(buf.String())
	os.Stdout = rescueStdout
}

/*
2023/02/09 17:56:50 Fake socks open port :1025
2023/02/09 17:56:50 Fake socks open port :1024
2023/02/09 17:56:51 Fake socks error addr  :1024 wrong ip
*/
