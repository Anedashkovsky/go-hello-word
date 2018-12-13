package config

import (
	"go-hello-word/internal/startupChecker"
	"io/ioutil"
	"log"
	"path/filepath"
)

type ConfigHelper struct {
	path string
	env  string
}

func (configHelper *ConfigHelper) Init(configName string, envName string) {
	envChecker := new(checker.EnvChecker)
	configHelper.env = envChecker.GetEnv(envName)
	configHelper.path = filepath.Join(configHelper.getConfigFolder(), configHelper.env, configName)
}

func (configHelper *ConfigHelper) GetConfig() []byte {
	config, error := ioutil.ReadFile(configHelper.path)

	if error != nil {
		log.Fatalln("Error while reading config", error)
	}

	return config
}

func (configHelper *ConfigHelper) getConfigFolder() string {
	path, error := filepath.Abs("../config")

	if error != nil {
		log.Fatalln("Error while create full filepath", error)
	}

	return path
}
