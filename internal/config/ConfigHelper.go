package config

import (
	"go-hello-word/internal/startupChecker"
	"io/ioutil"
	"log"
	"path/filepath"
)

// Helper is a struct to stor anv and path to config to read it
type Helper struct {
	path string
	env  string
}

// Init inititalizes struct with given configname and env
func (configHelper *Helper) Init(configName string, envName string) {
	envChecker := new(checker.EnvChecker)
	configHelper.env = envChecker.GetEnv(envName)
	configHelper.path = filepath.Join(configHelper.getConfigFolder(), configHelper.env, configName)
}

// GetConfig read file from disk and return byte array to unmarshall it
func (configHelper *Helper) GetConfig() []byte {
	config, error := ioutil.ReadFile(configHelper.path)

	if error != nil {
		log.Fatalln("Error while reading config", error)
	}

	return config
}

func (configHelper *Helper) getConfigFolder() string {
	path, error := filepath.Abs("../config")

	if error != nil {
		log.Fatalln("Error while create full filepath", error)
	}

	return path
}
