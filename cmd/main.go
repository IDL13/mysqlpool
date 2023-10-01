package main

import (
	"fmt"
	"mysqlpool/pkg/config"
)

func main() {
	conf := config.New().GetReplicationConfig()

	fmt.Println(conf.Tables[0])
}
