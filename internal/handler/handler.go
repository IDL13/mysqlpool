package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
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
	status := h.Balancer.GetHealth()
	for name, s := range status {
		resp.Write([]byte(name + "\n"))
		resp.Write([]byte(boolString(s)))
	}
}

func (h *Handler) ReadHandler(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		reqJson := req.Body

		var q queryStruct

		json.NewDecoder(reqJson).Decode(&q)

		conf, err := h.Balancer.Connections.GetConnection()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
			os.Exit(1)
		}

		if conf.Slave1.Count > conf.Slave2.Count {
			h.wg.Add(1)

			go func() {
				_, err = conf.Slave1.Slave1Compound.Exec(q.SqlQuery, q.Args)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Fale Exec request: %v\n", err)
					os.Exit(1)
				}

				atomic.AddUint64(&conf.Slave1.Count, 1)

				h.wg.Done()
			}()

			h.wg.Wait()
		} else {
			h.wg.Add(1)

			go func() {
				_, err = conf.Slave2.Slave2Compound.Exec(q.SqlQuery, q.Args)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Fale Exec request: %v\n", err)
					os.Exit(1)
				}

				atomic.AddUint64(&conf.Slave2.Count, 1)

				h.wg.Done()
			}()

			h.wg.Wait()
		}

	} else {
		resp.Write([]byte("This url only handles POST requests"))
	}
}

func (h *Handler) InsertHandler(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		reqJson := req.Body

		var q queryStruct

		json.NewDecoder(reqJson).Decode(&q)

		conf, err := h.Balancer.Connections.GetConnection()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
			os.Exit(1)
		}

		_, err = conf.Main.MainCompound.Exec(q.SqlQuery, q.Args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fale Exec request: %v\n", err)
			os.Exit(1)
		}

	} else {
		resp.Write([]byte("This url only handles POST requests"))
	}
}

func (h *Handler) MakeMigrationHandler(resp http.ResponseWriter, req *http.Request) {

}

func (h *Handler) MigrateHandler(resp http.ResponseWriter, req *http.Request) {

}
