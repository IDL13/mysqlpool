package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func New() ConfigurationInterface {
	configuration := &configurationStruct{}
	return configuration
}

type configurationStruct struct {
	dbConfig          *dbConfig
	replicationConfig *replicationConfig
}

type dbConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
}

type replicationConfig struct {
	Name string `yaml:"name"`
	Copy int    `yaml:"copy"`
}

type ConfigurationInterface interface {
	DbInterface
	ReplicationInterface
}

type DbInterface interface {
	GetDbConfig() *dbConfig
}

type ReplicationInterface interface {
	GetReplicationConfig() *replicationConfig
}

func (db *configurationStruct) GetDbConfig() *dbConfig {
	info, err := os.ReadFile("./config.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading data from conf file: %v\n", err)
		os.Exit(1)
	}
	yaml.Unmarshal(info, db.dbConfig)
	return db.dbConfig
}

func (replication *configurationStruct) GetReplicationConfig() *replicationConfig {
	info, err := os.ReadFile("./config.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading data from conf file: %v\n", err)
		os.Exit(1)
	}
	yaml.Unmarshal(info, replication.replicationConfig)
	return replication.replicationConfig
}
