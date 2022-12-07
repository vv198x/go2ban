package main

import (
	"go2ban/apiForgRPC"
	"go2ban/cmd/commandLine"
	"go2ban/cmd/firewall"
	"go2ban/pkg/config"
	"go2ban/pkg/logger"
)

func main() {
	config.Load()
	logger.Start()
	commandLine.Run()
	firewall.WorkerStart(config.Get().Flags.RunAsDaemon)
	apiForgRPC.Start("tcp", ":2048", config.Get().Flags.RunAsDaemon)

}
