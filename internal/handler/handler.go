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

	for name, s := range status {
		resp.Write([]byte(name + "\t"))
		resp.Write([]byte(boolString(s) + "\n"))
	}
}

func (h *Handler) ReadHandler(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {

		row, err := h.router.Redirection("read").Query("select * from name")
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
	_, err := h.router.Redirection("write").Exec(`insert into name(id, name) value(2, "Oleg")`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fale Exec request: %v\n", err)
		os.Exit(1)
	}
	h.router.Migrate(`insert into name(id, name) value(2, "oleg")`)
}

func (h *Handler) MigrateHandler(resp http.ResponseWriter, req *http.Request) {

}
