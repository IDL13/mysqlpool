package main

import (
	"fmt"
	"mysqlpool/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//signals
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	stop := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	a := server.New()
	go a.Run(stop)

	select {
	case stopR := <-done:
		fmt.Println(stopR)
		os.Exit(1)
	case stopL := <-stop:
		fmt.Println(stopL)
		os.Exit(1)
	}
}
