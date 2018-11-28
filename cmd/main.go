package main

import (
	"fmt"
	"go-hello-word/internal/config"
)

func main() {
	var serverConfig = new(config.ServerConfig)
	serverConfig.Init()
	fmt.Println("ServerConfig=", serverConfig.GetHost(), serverConfig.GetPort())
}
