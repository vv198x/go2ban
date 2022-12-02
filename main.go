package main

import (
	"microservice2ban/apiForgRPC"
	"microservice2ban/pkg/logger"
)

func main() {
	logger.Start()
	apiForgRPC.Start("tcp", ":2048")
}
