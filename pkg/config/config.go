package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func (c configuration) GetDbConfig() *configuration {

	// Create new struct *configuration
	conf := new()

	info, err := os.ReadFile("./conf.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading data from conf file: %v\n", err)
		os.Exit(1)
	}

	// unmarshaling date in struct
	yaml.Unmarshal(info, conf.DbConfigMain)
	yaml.Unmarshal(info, conf.DbConfigSlave1)
	yaml.Unmarshal(info, conf.DbConfigSlave2)

	return conf
}

//func (c configuration) GetReplicationConfig() *replicationConfig {
//	conf := &replicationConfig{}
//
//	info, err := os.ReadFile("./conf.yaml")
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Error reading data from conf file: %v\n", err)
//		os.Exit(1)
//	}
//
//	yaml.Unmarshal(info, conf)
//
//	return conf
//}
