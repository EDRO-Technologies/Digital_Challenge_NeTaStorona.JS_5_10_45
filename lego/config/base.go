package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port           string
	Environment    string
	Host           string
	ServiceName    string
	Version        string
	Model          string
	ModelToken     string
	EventGathering string
}

type confVars struct {
	missing   []string //name of the mandatory environment variable that are missing
	malformed []string //errors describing malformed environment varibale values
}

var Conf *Config

func New() (*Config, error) {
	vars := &confVars{}

	port := vars.mandatoryInt("PORT")
	environment := vars.mandatory("ENVIRONMENT")
	host := vars.mandatory("HOST")
	serviceName := vars.mandatory("SERVICE_NAME")
	version := vars.mandatory("VERSION")
	model := vars.mandatory("MODEL_URL")
	modelToken := vars.mandatory("MODEL_TOKEN")
	//eventGathering := vars.mandatory("EVENT_GATHERING")

	if err := vars.Error(); err != nil {
		return nil, fmt.Errorf("error loading configuration: %w", err)
	}

	config := &Config{
		Port:        fmt.Sprintf(":%d", port),
		Environment: environment,
		Host:        host,
		ServiceName: serviceName,
		Version:     version,
		Model:       model,
		ModelToken:  modelToken,
		//EventGathering: eventGathering,
	}

	Conf = config

	return config, nil
}

func (vars *confVars) mandatory(key string) string {
	value := os.Getenv(key)
	if value == "" {
		vars.missing = append(vars.missing, key)
	}
	return value
}

func (vars *confVars) mandatoryInt(key string) int {
	value := os.Getenv(key)
	if value == "" {
		vars.missing = append(vars.missing, key)
		return 0
	}

	valueInt, err := strconv.Atoi(value)

	if err != nil {
		vars.malformed = append(vars.malformed, key)
		return 0
	}

	return valueInt
}

//func (vars *confVars) mandatoryDuration(key string) time.Duration {
//	value := os.Getenv(key)
//	if value == "" {
//		vars.missing = append(vars.missing, key)
//		return 0
//	}
//
//	valueDuration, err := time.ParseDuration(value)
//
//	if err != nil {
//		vars.malformed = append(vars.malformed, key)
//		return 0
//	}
//
//	return valueDuration
//}

//func (vars *confVars) mandatoryBool(key string) bool {
//	value := os.Getenv(key)
//	if value == "" {
//		vars.missing = append(vars.missing, key)
//		return false
//	}
//
//	valueBool, err := strconv.ParseBool(value)
//
//	if err != nil {
//		vars.malformed = append(vars.malformed, key)
//		return false
//	}
//
//	return valueBool
//}

func (vars confVars) Error() error {
	if len(vars.missing) > 0 {
		return fmt.Errorf("missing mandatory configurations: %s", strings.Join(vars.missing, ", "))
	}

	if len(vars.malformed) > 0 {
		return fmt.Errorf("malformed configurations: %s", strings.Join(vars.malformed, "; "))
	}
	return nil
}
