package server

import (
	"fmt"
	"mysqlpool/internal/handler"
	"net/http"
	"os"
	"time"
)

type Server struct {
	s   *http.Server
	h   handler.Handler
	mux *http.ServeMux
}

func New() *Server {
	s := &Server{
		s: &http.Server{
			Addr:           ":3333",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		h:   handler.New(),
		mux: http.NewServeMux(),
	}

	s.s.Handler = s.mux

	s.mux.HandleFunc("/", s.h.StartServer)
	s.mux.HandleFunc("/probe", s.h.HealthinessProbe)
	s.mux.HandleFunc("/read", s.h.ReadHandler)
	s.mux.HandleFunc("/insert", s.h.InsertHandler)
	s.mux.HandleFunc("/create_test", s.h.TestHandler)

	return s
}

func (s *Server) Run(stop chan bool) {
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
