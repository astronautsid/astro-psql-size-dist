package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConfig(configFile string) (Config, error) {
	var config Config

	file, err := os.Open(configFile)
	if err != nil {
		return config, err
	}

	defer file.Close()

	if file != nil {
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&config); err != nil {
			return config, err
		}
	}

	return config, nil
}
