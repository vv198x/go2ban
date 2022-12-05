package main

import (
	"fmt"
	"go2ban/apiForgRPC"
	"go2ban/cmd/addFirewall"
	"go2ban/pkg/config"
	"go2ban/pkg/logger"
)

func main() {
	config.Load()
	logger.Start()
	apiForgRPC.Start("tcp", ":2048", config.Get().RunAsDaemon)
	fmt.Println(addFirewall.FirewallUnBlockAll())
}
