package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
	} `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	var config Config
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
