package router

import (
	"database/sql"
	"fmt"
	"mysqlpool/pkg/mysqlconn"
	"os"
	"sync"
)

func New() *Router {
	return &Router{
		Compound: mysqlconn.New(),
	}
}

type Router struct {
	Compound *mysqlconn.Compound
}

func (r *Router) Redirection(flag string) *sql.DB {
	conf, err := r.Compound.GetConnection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
		os.Exit(1)
	}

	switch flag {
	case "write":
		return conf.Main.MainCompound
	case "read":
		status := mysqlconn.GetHealth(conf)
		if status["slave1"] == true {
			return conf.Slave1.Slave1Compound
		} else {
			return conf.Slave2.Slave2Compound
		}
	default:
		return conf.Main.MainCompound
	}
}

func (r *Router) Migrate(q string) error {
	var wg sync.WaitGroup

	conf, err := r.Compound.GetConnection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
		os.Exit(1)
	}

	wg.Add(2)

	go func() {
		_, err = conf.Slave1.Slave1Compound.Exec(q)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to Exex {router 51}: %v\n", err)
			os.Exit(1)
		}

		wg.Done()
	}()

	go func() {
		_, err = conf.Slave2.Slave2Compound.Exec(q)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to Exex {router 41}: %v\n", err)
			os.Exit(1)
		}

		wg.Done()
	}()

	wg.Wait()

	return nil
}
