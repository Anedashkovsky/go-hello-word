package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ConfigHelper struct {
	path string
	env  string
}

func (configHelper *ConfigHelper) Init(configName string, envName string) {
	configHelper.env = os.Getenv(envName)
	configHelper.path = filepath.Join(configHelper.getConfigFolder(), configHelper.env, configName)
}

func (configHelper *ConfigHelper) GetConfig() []byte {
	config, error := ioutil.ReadFile(configHelper.path)

	if error != nil {
		fmt.Println("Error while reading config", error)
	}

	return config
}

func (configHelper *ConfigHelper) getConfigFolder() string {
	path, error := filepath.Abs("../config")

	if error != nil {
		fmt.Println("Error while create full filepath")
	}

	return path
}
