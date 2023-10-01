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

		// Query string
		q := fmt.Sprintf(`CREATE DATABASE %s_%d`, conf.Name, conf.Copy)

		_, err := client.Query(q)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while adding data: %v\n", err)
			os.Exit(1)
		}
	}
}
