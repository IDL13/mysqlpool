package balancer

import (
	"errors"
	"fmt"
	"mysqlpool/pkg/mysqlconn"
	"os"
)

func New() *Balancer {
	return &Balancer{
		Connections: mysqlconn.New(),
	}
}

type Balancer struct {
	Connections *mysqlconn.Compound
}

func (b *Balancer) handler(sqlQuery, request, mod string, args ...any) (interface{}, error) {
	switch request {
	case "Exec":
		conf, err := b.Connections.GetConnection()
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
		conf, err := b.Connections.GetConnection()
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
		conf, err := b.Connections.GetConnection()
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

func (b *Balancer) GetHealth() map[string]bool {
	status := make(map[string]bool)

	status["main"] = b.Connections.Main.HealthinessProbe
	status["slave1"] = b.Connections.Slave1.HealthinessProbe
	status["slave2"] = b.Connections.Slave2.HealthinessProbe

	return status
}
