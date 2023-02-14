package localService

import (
	"context"
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/vv198x/go2ban/cmd/firewall"
	"github.com/vv198x/go2ban/config"
	"os"
	"testing"
	"time"
)

func TestWorkerStart(t *testing.T) {
	// Set up test data
	firewall.ExportFirewall = &firewall.Mock{}
	mockCtx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()
	//service1 := config.Service{On: true, Name: "service1", LogFile: "/etc/passwd", Regxp: ".*error.*"}
	service2 := config.Service{On: true, Name: "service2", LogFile: "docker", Regxp: ".*panic.*"}
	service3 := config.Service{On: true, LogFile: "test.log", Name: "docker", Regxp: "pattern2"}
	service4 := config.Service{On: false, LogFile: "test2.log", Name: "Service 3", Regxp: "pattern3"}

	// Reset the flags before the test
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	config.Get().ServiceCheckMinutes = 2
	config.Get().LogDir = "."

	t.Run("Test the function with the RunAsDaemon flag false", func(t *testing.T) {
		WorkerStart(context.TODO(), false, []config.Service{}, nil)
		// If it does not wait without a mock context, then everything is OK
	})

	t.Run("Test the function with a mock context and off log file", func(t *testing.T) {
		WorkerStart(mockCtx, true, []config.Service{service4}, nil)

	})

	t.Run("Test the function with a mock context and docker logfile", func(t *testing.T) {
		WorkerStart(mockCtx, true, []config.Service{service2, service3}, nil)
	})

	t.Run("Check that the map file was saved correctly", func(t *testing.T) {
		err := os.Remove("endBytesMap")
		assert.NoError(t, err)
	})
}
