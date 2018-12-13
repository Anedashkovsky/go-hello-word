package config

import (
	"encoding/json"
	"log"
)

// Server struct to store unmarshalled config data
type Server struct {
	rawConfig rawServerConfig
}

type rawServerConfig struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

const configName = "server.json"
const envName = "ENV"

// Init create dependent instances of structs, take config file from disk and unmarshall it
func (serverConfig *Server) Init() {
	configHelper := new(Helper)
	configHelper.Init(configName, envName)

	error := json.Unmarshal(configHelper.GetConfig(), &serverConfig.rawConfig)

	if error != nil {
		log.Fatalln("Error unmarshalling config", error)
	}
}

// GetPort return the port for server
func (serverConfig *Server) GetPort() string {
	return serverConfig.rawConfig.Port
}

// GetHost return host for server
func (serverConfig *Server) GetHost() string {
	return serverConfig.rawConfig.Host
}
