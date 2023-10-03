package replication

import (
	"database/sql"
	"fmt"
	"mysqlpool/pkg/config"
	"mysqlpool/pkg/mysqlconn"
	"os"
	"sync"
)

func createMainDb(client *sql.DB, name string) {

	// Query string (create main db)
	q := fmt.Sprintf(`CREATE DATABASE %s`, name)

	// Create main database
	_, err := client.Exec(q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while creating replication databases: %v\n", err)
		os.Exit(1)
	}
}

func createTables(client *sql.DB, tableString string) {
	_, err := client.Exec(tableString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while creating tables: %v\n", err)
		os.Exit(1)
	}
}

func createReplication() {
	var wg sync.WaitGroup

	// Create replication config
	conf := config.New().GetReplicationConfig()

	// Db client
	client, err := mysqlconn.NewClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading data from conf file: %v\n", err)
		os.Exit(1)
	}

	// Create main db
	go createMainDb(client, conf.Name)

	wg.Add(conf.Copy)

	//Replication loop
	for i := 0; i < conf.Copy; i++ {
		go func() {

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

			wg.Done()
		}()

		wg.Wait()
	}
}
