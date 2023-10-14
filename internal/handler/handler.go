package handler

import (
	"fmt"
	"mysqlpool/pkg/mysqlconn"
	"net/http"
	"os"
)

func boolString(b bool) string {
	if b == true {
		return "true"
	} else {
		return "false"
	}
}

//test handlers for example

func (h *Handler) StartServer(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("SERVER STARTED"))
}

func (h *Handler) HealthinessProbe(resp http.ResponseWriter, req *http.Request) {
	conf, err := h.Compound.GetConnection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
		os.Exit(1)
	}

	status := mysqlconn.GetHealth(conf)

	m, s1, s2 := h.counter.LoadCounter()

	resp.Write([]byte("main connections" + "\t"))
	resp.Write([]byte(h.counter.ConvertOnString(m) + "\n"))

	resp.Write([]byte("slave1 connections" + "\t"))
	resp.Write([]byte(h.counter.ConvertOnString(s1) + "\n"))

	resp.Write([]byte("Slave2 connections" + "\t"))
	resp.Write([]byte(h.counter.ConvertOnString(s2) + "\n"))

	for name, s := range status {
		resp.Write([]byte(name + "\t"))
		resp.Write([]byte(boolString(s) + "\n"))
	}
}

func (h *Handler) ReadHandler(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {

		row, err := h.router.Redirection("r").Query("select * from name")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fale Exec request: %v\n", err)
			os.Exit(1)
		}
		for row.Next() {
			type Name struct {
				id   int
				name string
			}

			var name Name

			err = row.Scan(&name.id, &name.name)
			fmt.Println(name.id, name.name)
		}
	}
}

func (h *Handler) InsertHandler(resp http.ResponseWriter, req *http.Request) {
	_, err := h.router.Redirection("w").Exec(`insert into testing(id, name, count) value(1, "Ilya", 1)`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fale Exec request: %v\n", err)
		os.Exit(1)
	}
	h.router.Migrate(`insert into testing(id, name, count) value(1, "Ilya", 1)`)
}

func (h *Handler) TestHandler(resp http.ResponseWriter, req *http.Request) {

}
