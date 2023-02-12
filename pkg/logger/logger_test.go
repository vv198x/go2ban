package logger

import (
	"github.com/vv198x/go2ban/config"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	config.Get().LogDir = "/tmp"

	// call Start and check if the log file has been created
	Start()
	logFilePath := filepath.Join("/tmp", time.Now().Format("06.01.02")+logExp)
	_, err := os.Stat(logFilePath)
	if err != nil {
		t.Fatalf("log file not created: %v", err)
	}

	// write some log messages and check if they are written to the log file
	log.Println("test log message 1")
	log.Println("test log message 2")
	file, openError := os.Open(logFilePath)
	if openError != nil {
		t.Fatalf("The log file could not be opened: %v", openError)
	}
	defer file.Close()

	stat, statError := file.Stat()
	if statError != nil {
		t.Fatalf("The log file stat could not be retrieved: %v", statError)
	}

	content := make([]byte, stat.Size())
	_, readError := file.Read(content)
	if readError != nil {
		t.Fatalf("The log file could not be read: %v", readError)
	}

	expected := "test log message 2\n"
	if !strings.Contains(string(content), expected) {
		t.Fatalf("unexpected log file content, expected: %q, got: %q", expected, string(content))
	}
}
