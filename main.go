package main

import (
	"microservice2ban/cmd/addFirewall"
	"microservice2ban/pkg/logger"
)

func main() {
	logger.Start()
	//apiForgRPC.Start("tcp", ":2048")
	addFirewall.BlockIP("1")
}
