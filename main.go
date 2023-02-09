package main

import (
	"github.com/vv198x/go2ban/api/REST"
	"github.com/vv198x/go2ban/api/gRPC"
	"github.com/vv198x/go2ban/cmd/commandLine"
	"github.com/vv198x/go2ban/cmd/fakeSocks"
	"github.com/vv198x/go2ban/cmd/firewall"
	"github.com/vv198x/go2ban/cmd/localService"
	"github.com/vv198x/go2ban/pkg/config"
	"github.com/vv198x/go2ban/pkg/logger"
	proFile "github.com/vv198x/go2ban/pkg/profile"
)

func main() {
	pprofEnd := proFile.Start("nope")
	config.Load()
	logger.Start()
	commandLine.Run()
	firewall.WorkerStart(config.Get().Flags.RunAsDaemon)
	fakeSocks.Listen(config.Get().FakeSocksPorts)
	REST.Start(config.Get().Flags.RunAsDaemon)
	gRPC.Start(config.Get().Flags.RunAsDaemon)
	localService.WorkerStart(config.Get().Services, pprofEnd)
}
