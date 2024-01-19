package utils

import (
	"gopkg.in/yaml.v2"
	"os"
	"retrognome/internal/types"
)

func LoadApplicationConfiguration() (types.ApplicationConfiguration, error) {

	file, err := os.Open("config.yaml")

	if err != nil {
		return types.ApplicationConfiguration{}, err
	}
	defer file.Close()

	var config types.ApplicationConfiguration
	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		return types.ApplicationConfiguration{}, err
	}

	return config, nil
}
