package config

import (
	"encoding/json"
	"log"
)

type ServerConfig struct {
	rawConfig rawServerConfig
}

type rawServerConfig struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

const CONFIG_NAME = "server.json"
const ENV_NAME = "ENV"

func (serverConfig *ServerConfig) Init() {
	configHelper := new(ConfigHelper)
	configHelper.Init(CONFIG_NAME, ENV_NAME)

	error := json.Unmarshal(configHelper.GetConfig(), &serverConfig.rawConfig)

	if error != nil {
		log.Fatalln("Error unmarshalling config", error)
	}
}

func (serverConfig *ServerConfig) GetPort() string {
	return serverConfig.rawConfig.Port
}

func (serverConfig *ServerConfig) GetHost() string {
	return serverConfig.rawConfig.Host
}
