package checker

import (
	"errors"
	"log"
	"os"
)

// EnvChecker Struct to make operations at env variables like read
type EnvChecker struct{}

// NewEnvChecker Create env checker
func NewEnvChecker() *EnvChecker {
	return &EnvChecker{}
}

// GetEnv Get value of env variable
func (envChecker *EnvChecker) GetEnv(envName string) string {
	env, error := envChecker.getEnv(envName)

	if error != nil {
		log.Fatal(error)
	}

	return env
}

func (envChecker *EnvChecker) getEnv(envName string) (string, error) {
	var variableNotSet error
	env := os.Getenv(envName)

	if env == "" {
		variableNotSet = errors.New("Environment variable " + envName + " must be set at startup")
	}

	return env, variableNotSet
}
