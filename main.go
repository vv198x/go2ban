package main

import (
	"context"
	"github.com/vv198x/go2ban/api/grpc"
	"github.com/vv198x/go2ban/api/rest"
	"github.com/vv198x/go2ban/cmd/abuseipdb"
	"github.com/vv198x/go2ban/cmd/commandLine"
	"github.com/vv198x/go2ban/cmd/fakeSocks"
	"github.com/vv198x/go2ban/cmd/firewall"
	"github.com/vv198x/go2ban/cmd/localService"
	"github.com/vv198x/go2ban/config"
	"github.com/vv198x/go2ban/pkg/logger"
	proFile "github.com/vv198x/go2ban/pkg/profile"
)

func main() {
	pprofEnd := proFile.Start("nope")
	config.ReadFlags()
	config.Load()
	logger.Start()
	firewall.Initialization(config.Get().Flags.RunAsDaemon)
	commandLine.Run()
	fakeSocks.Listen(config.Get().TrapPorts)
	rest.Start(config.Get().Flags.RunAsDaemon)
	grpc.Start(config.Get().Flags.RunAsDaemon)
	abuseipdb.Scheduler(config.Get().AbuseipdbApiKey)
	localService.WorkerStart(
		context.TODO(),
		config.Get().Flags.RunAsDaemon,
		config.Get().Services,
		pprofEnd) // To stop profiling
}
