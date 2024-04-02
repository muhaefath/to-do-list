package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port     int    `yaml:"port"`
	LogLevel string `yaml:"log_level"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func LoadConfig(configFile string) (Config, error) {
	var config Config

	// Read YAML file
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("Error reading YAML file: %v\n", err)
		return config, err
	}

	// Unmarshal YAML data into struct
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Printf("Error unmarshalling YAML: %v\n", err)
		return config, err
	}

	return config, nil
}
