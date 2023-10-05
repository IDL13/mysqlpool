package handler

import (
	"mysqlpool/internal/balancer"
	"net/http"
)

func New() Handler {
	return Handler{
		balancer: balancer.New(),
	}
}

type Handler struct {
	balancer *balancer.Balancer
}

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
	status := h.balancer.GetHealth()
	for name, s := range status {
		resp.Write([]byte(name + "\n"))
		resp.Write([]byte(boolString(s)))
	}
}
