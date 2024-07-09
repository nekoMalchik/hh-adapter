package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	Server struct {
		Port string
	}
	DB struct {
		Uri    string
		DbName string
	}
	HH struct {
		APIKey string
	}
}

func LoadConfig() (*Config, error) {
	cnf := &Config{}
	cnfPath := filepath.Join("internal", "config", "config.yml")
	yamlFile, err := os.ReadFile(cnfPath)

	if err != nil {
		fmt.Println("Error while reading yml file: ", err)
	}

	err = yaml.Unmarshal(yamlFile, cnf)

	if err != nil {
		fmt.Println("Unmarshal error: ", err)
	}

	return cnf, nil
}
