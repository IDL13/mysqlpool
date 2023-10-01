package replication

import (
	"fmt"
	"mysqlpool/pkg/config"
	"mysqlpool/pkg/mysqlconn"
	"os"
)

func createReplication() {
	// Create replication config
	conf := config.New().GetReplicationConfig()

	// Db client
	client, err := mysqlconn.NewClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading data from conf file: %v\n", err)
		os.Exit(1)
	}

	//Replication loop
	for i := 0; i < conf.Copy; i++ {

		// Query string (create replication db)
		q := fmt.Sprintf(`CREATE DATABASE %s_%d`, conf.Name, conf.Copy)

		_, err := client.Exec(q)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while creating replication databases: %v\n", err)
			os.Exit(1)
		}

		// Query string (change databases)
		q = fmt.Sprintf(`USE %s_%d`, conf.Name, conf.Copy)

		_, err = client.Exec(q)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while changing database: %v\n", err)
			os.Exit(1)
		}

		// Creating tables loop
		for _, j := range conf.Tables {
			_, err = client.Exec(j)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error while creating tables: %v\n", err)
				os.Exit(1)
			}
		}
	}
}
