package config

func New() Configuration {
	return &configuration{}
}

func new() *configuration {
	return &configuration{
		DbConfigMain:   &dbConfigMain{},
		DbConfigSlave1: &dbConfigSlave1{},
		DbConfigSlave2: &dbConfigSlave2{},
	}
}

type Configuration interface {
	GetDbConfig() *configuration
	GetReplicationConfig() *replicationConfig
}

type configuration struct {
	DbConfigMain   *dbConfigMain
	DbConfigSlave1 *dbConfigSlave1
	DbConfigSlave2 *dbConfigSlave2
}

type dbConfigMain struct {
	User     string `yaml:"user_main"`
	Password string `yaml:"password_main"`
	Host     string `yaml:"host_main"`
	Port     string `yaml:"port_main"`
	Db       string `yaml:"db_main"`
}

type dbConfigSlave1 struct {
	User     string `yaml:"user_slave_1"`
	Password string `yaml:"password_slave_1"`
	Host     string `yaml:"host_main_slave_1"`
	Port     string `yaml:"port_main_slave_1"`
	Db       string `yaml:"db_main_slave_1"`
}

type dbConfigSlave2 struct {
	User     string `yaml:"user_slave_2"`
	Password string `yaml:"password_slave_2"`
	Host     string `yaml:"host_main_slave_2"`
	Port     string `yaml:"port_main_slave_2"`
	Db       string `yaml:"db_main_slave_2"`
}

type replicationConfig struct {
	Name   string   `yaml:"name"`
	Copy   int      `yaml:"copy"`
	Tables []string `yaml:"tables_string"`
}
