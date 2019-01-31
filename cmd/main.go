package main

import (
	"go-hello-word/internal/config"
	"go-hello-word/internal/server"
)

func main() {
	var serverConfig = new(config.Server)
	serverConfig.Init()
	var server = new(server.Server)
	server.Init(serverConfig)
	server.Start()
}
