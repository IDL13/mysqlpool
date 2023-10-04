package balancer

import (
	"errors"
	"fmt"
	"mysqlpool/pkg/mysqlconn"
	"os"
)

type Balancer struct {
	connections mysqlconn.Compound
}

func (b *Balancer) handler(sqlQuery, request, mod string, args ...any) (interface{}, error) {
	switch request {
	case "Exec":
		conf, err := mysqlconn.New().GetConnection()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
			os.Exit(1)
		}
		res, err := conf.Main.MainCompound.Exec(sqlQuery, args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fale Exec request: %v\n", err)
			os.Exit(1)
		}

		return res, nil

	case "Query":
		conf, err := mysqlconn.New().GetConnection()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
			os.Exit(1)
		}
		row, err := conf.Slave1.Slave1Compound.Query(sqlQuery, args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fale Query request: %v\n", err)
			os.Exit(1)
		}

		return row, nil

	case "QueryRow":
		conf, err := mysqlconn.New().GetConnection()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
			os.Exit(1)
		}
		row := conf.Slave1.Slave1Compound.QueryRow(sqlQuery, args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fale Query request: %v\n", err)
			os.Exit(1)
		}

		return row, nil
	default:
		err := errors.New("Unknowen argyments")
		return err, err
	}
}
