package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func New() Configuration {
	return &configuration{}
}

type configuration struct {
}

type dbConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
}

type replicationConfig struct {
	Name   string   `yaml:"name"`
	Copy   int      `yaml:"copy"`
	Tables []string `yaml:"tables_string"`
}

type Configuration interface {
	GetDbConfig() *dbConfig
	GetReplicationConfig() *replicationConfig
}

func (c configuration) GetDbConfig() *dbConfig {
	conf := &dbConfig{}
	info, err := os.ReadFile("./conf.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading data from conf file: %v\n", err)
		os.Exit(1)
	}
	yaml.Unmarshal(info, conf)
	return conf
}

func (c configuration) GetReplicationConfig() *replicationConfig {
	conf := &replicationConfig{}
	info, err := os.ReadFile("./conf.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading data from conf file: %v\n", err)
		os.Exit(1)
	}
	yaml.Unmarshal(info, conf)
	return conf
}
