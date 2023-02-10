package main

import (
	"github.com/vv198x/go2ban/api/grpc"
	"github.com/vv198x/go2ban/api/rest"
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
	commandLine.Run()
	firewall.Initialization(config.Get().Flags.RunAsDaemon)
	fakeSocks.Listen(config.Get().FakeSocksPorts)
	rest.Start(config.Get().Flags.RunAsDaemon)
	grpc.Start(config.Get().Flags.RunAsDaemon)
	localService.WorkerStart(config.Get().Services, pprofEnd)
}
