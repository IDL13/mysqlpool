package mysqlconn

import (
	"database/sql"
	"fmt"
	"mysqlpool/pkg/config"
	"os"
)

func NewClient() (conn *sql.DB, err error) {
	// Create db configuration
	conf := config.New().GetDbConfig()

	q := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Db)

	conn, err = sql.Open("mysql", q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
		os.Exit(1)
	}

	return conn, nil
}
