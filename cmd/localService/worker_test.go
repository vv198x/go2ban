package localService

import (
	"flag"
	"github.com/vv198x/go2ban/cmd/firewall"
	"github.com/vv198x/go2ban/config"
	"os"
	"testing"
)

func TestWorkerStart1(t *testing.T) {
	firewall.ExportFirewall = &firewall.Mock{}

	// Reset the flags before the test
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	//Check not run
	config.Get().Flags.RunAsDaemon = false
	config.Get().ServiceCheckMinutes = 2

	// Test the function with a slice of services проверить выключенные
	service1 := config.Service{On: true, LogFile: "test.log"}
	service2 := config.Service{On: true, LogFile: "docker"}
	services := []config.Service{service1, service2}

	type args struct {
		services []config.Service
		pprofEnd interface{ Stop() }
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test worker start with RunAsDaemon flag false",
			args: args{
				services: []config.Service{},
				pprofEnd: nil,
			},
		},
		{
			name: "Test worker start with RunAsDaemon flag true and one service",
			args: args{
				services: services,
				pprofEnd: nil,
			},
		},
		{
			name: "Test worker start with RunAsDaemon flag true and two services, one of which is docker",
			args: args{
				services: []config.Service{{On: true, LogFile: "test.log"}, {On: true, LogFile: "docker"}},
				pprofEnd: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WorkerStart(tt.args.services, tt.args.pprofEnd)
		})
	}
}
