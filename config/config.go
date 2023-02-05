package config

import (
	"encoding/json"
	"io"
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

func ReadRule(ruleFile string) (Rule, error) {
	var rule Rule

	file, err := os.Open(ruleFile)
	if err != nil {
		return rule, err
	}

	defer file.Close()

	if file != nil {
		byteValue, err := io.ReadAll(file)
		if err != nil {
			return rule, err
		}
		err = json.Unmarshal(byteValue, &rule)
		if err != nil {
			return rule, err
		}
	}

	return rule, nil
}
