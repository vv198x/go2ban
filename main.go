package main

import (
	"go2ban/apiForgRPC"
	"go2ban/pkg/config"
	"go2ban/pkg/logger"
)

func main() {
	config.Load()
	logger.Start()
	apiForgRPC.Start("tcp", ":2048", config.Get().RunAsDaemon)
}
