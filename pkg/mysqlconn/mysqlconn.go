package mysqlconn

import (
	"database/sql"
	"fmt"
	"mysqlpool/pkg/config"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func New() *Compound {
	return &Compound{}
}

type Compound struct {
	MainCompound   *sql.DB
	Slave1Compound *sql.DB
	Slave2Compound *sql.DB
}

func (c *Compound) GetConnection() (conn *Compound, err error) {

	// Create db configuration
	conf := config.New().GetDbConfig()

	mainQ := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.DbConfigMain.User,
		conf.DbConfigMain.Password,
		conf.DbConfigMain.Host,
		conf.DbConfigMain.Port,
		conf.DbConfigMain.Db,
	)

	slave1Q := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.DbConfigSlave1.User,
		conf.DbConfigSlave1.Password,
		conf.DbConfigSlave1.Host,
		conf.DbConfigSlave1.Port,
		conf.DbConfigSlave1.Db,
	)

	slave2Q := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.DbConfigSlave2.User,
		conf.DbConfigSlave2.Password,
		conf.DbConfigSlave2.Host,
		conf.DbConfigSlave2.Port,
		conf.DbConfigSlave2.Db,
	)

	c.MainCompound, err = sql.Open("mysql", mainQ)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
		os.Exit(1)
	}

	c.Slave1Compound, err = sql.Open("mysql", slave1Q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
		os.Exit(1)
	}

	c.Slave2Compound, err = sql.Open("mysql", slave2Q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
		os.Exit(1)
	}

	return c, nil
}
