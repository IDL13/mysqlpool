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
	Main          mainServer
	Slave1        slave1Server
	Slave2        slave2Server
	RowFunctional []string
}

type mainServer struct {
	MainCompound     *sql.DB
	HealthinessProbe bool
}

type slave1Server struct {
	Slave1Compound   *sql.DB
	HealthinessProbe bool
	Count            uint64
}

type slave2Server struct {
	Slave2Compound   *sql.DB
	HealthinessProbe bool
	Count            uint64
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

	c.Main.MainCompound, err = sql.Open("mysql", mainQ)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
		os.Exit(1)
	}

	c.Slave1.Slave1Compound, err = sql.Open("mysql", slave1Q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
		os.Exit(1)
	}

	c.Slave1.Count = 1

	c.Slave2.Slave2Compound, err = sql.Open("mysql", slave2Q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
		os.Exit(1)
	}

	c.Slave2.Count = 1

	return c, nil
}

func GetHealth(c *Compound) map[string]bool {
	connection, err := c.GetConnection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unbale to connection to database 97 string: %v\n", err)
		os.Exit(1)
	}

	status := make(map[string]bool)

	_, err = connection.Main.MainCompound.Exec(`SHOW DATABASES`)
	if err != nil {
		status["main"] = false
	} else {
		status["main"] = true
	}

	_, err = connection.Slave1.Slave1Compound.Exec(`SHOW DATABASES`)
	if err != nil {
		status["slave1"] = false
	} else {
		status["slave1"] = true
	}

	_, err = connection.Slave2.Slave2Compound.Exec(`SHOW DATABASES`)
	if err != nil {
		status["slave2"] = false
	} else {
		status["slave2"] = true
	}

	return status
}
