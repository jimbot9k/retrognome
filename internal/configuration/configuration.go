package configuration

import (
	"gopkg.in/yaml.v2"
	"os"
)

type ApplicationConfiguration struct {
	AppName string `yaml:"app_name"`
	Port    int    `yaml:"port"`
}

func LoadApplicationConfiguration() (ApplicationConfiguration, error) {

	file, err := os.Open("config.yaml")

	if err != nil {
		return ApplicationConfiguration{}, err
	}
	defer file.Close()

	var config ApplicationConfiguration
	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		return ApplicationConfiguration{}, err
	}

	return config, nil
}
