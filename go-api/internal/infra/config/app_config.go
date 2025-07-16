// internal/infra/config/config.go
package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Application struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"application"`
	API2 struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Endpoint string `yaml:"endpoint"`
	} `yaml:"api2"`
}

func LoadConfig() *Config {
	var config Config

	// Read the YAML file
	file, err := os.ReadFile("config/application.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Parse YAML
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	return &config
}