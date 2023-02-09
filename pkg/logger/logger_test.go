package logger

import (
	"github.com/vv198x/go2ban/pkg/config"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	defer func() {
		log.SetOutput(os.Stderr)
		if logFile != nil {
			logFile.Close()
			os.Remove(logFile.Name())
		}
	}()

	// set up a temp directory for the log file
	tempDir, err := ioutil.TempDir("/tmp", "")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)
	config.Get().LogDir = tempDir

	// call Start and check if the log file has been created
	Start()
	logFilePath := filepath.Join(tempDir, time.Now().Format("06.01.02")+logExp)
	_, err = os.Stat(logFilePath)
	if err != nil {
		t.Fatalf("log file not created: %v", err)
	}

	// write some log messages and check if they are written to the log file
	log.Println("test log message 1")
	log.Println("test log message 2")
	content, err := ioutil.ReadFile(logFilePath)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}
	expected := "test log message 2\n"
	if !strings.Contains(string(content), expected) {
		t.Fatalf("unexpected log file content, expected: %q, got: %q", expected, string(content))
	}
}
