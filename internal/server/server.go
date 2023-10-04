package server

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Server struct {
	s   *http.Server
	h   hadler.Handler
	mux *http.ServeMux
}

func New() *Server {
	s := Server{
		s: &http.Server{
			Addr:           ":3333",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		h:   hadnler.New(),
		mux: http.NewServeMux(),
	}

	s.mux.HandleFunc("/", s.h.StartServer)
	s.mux.HandleFunc("probe", s.h.HealthinessProbe)

	return s
}

func (s *Server) Run() {
	fmt.Println(`
╔══╗╔══╗─╔╗───╔╗╔══╗
╚╗╔╝║╔╗╚╗║║──╔╝║╚═╗║
─║║─║║╚╗║║║──╚╗║╔═╝║
─║║─║║─║║║║───║║╚═╗║
╔╝╚╗║╚═╝║║╚═╗─║║╔═╝║
╚══╝╚═══╝╚══╝─╚╝╚══╝
	`)
	fmt.Println("[SERVER STARTED]")
	fmt.Println("http://127.0.0.1:3333")
	if err := s.s.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}