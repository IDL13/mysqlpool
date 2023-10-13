package handler

import (
	"mysqlpool/internal/counter"
	"mysqlpool/internal/router"
	"mysqlpool/pkg/mysqlconn"
	"sync"
)

func New() Handler {
	return Handler{
		Compound: mysqlconn.New(),
		router:   router.New(),
		counter:  counter.New(),
	}
}

type Handler struct {
	Compound    *mysqlconn.Compound
	QueryStruct *queryStruct
	wg          sync.WaitGroup
	router      *router.Router
	counter     *counter.Counter
}

type queryStruct struct {
	SqlQuery string   `json:"sqlQuery"`
	Args     []string `json:"args"`
}
