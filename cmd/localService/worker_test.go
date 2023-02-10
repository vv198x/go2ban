package localService

import (
	"flag"
	"github.com/vv198x/go2ban/config"
	"os"
	"testing"
)

func TestWorkerStart(t *testing.T) {
	// Reset the flags before the test
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	//Check not run
	config.Get().Flags.RunAsDaemon = false
	config.Get().ServiceCheckMinutes = 2

	var services []config.Service
	var pprofEnd interface{ Stop() }

	// Test the function with an empty slice of services
	WorkerStart(services, pprofEnd)

	// Test the function with a slice of services проверить выключенные
	service1 := config.Service{On: true, LogFile: "test.log"}
	service2 := config.Service{On: true, LogFile: "docker"}
	services = []config.Service{service1, service2}

	WorkerStart(services, pprofEnd)
}
