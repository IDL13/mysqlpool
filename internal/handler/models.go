package handler

import (
	"mysqlpool/internal/balancer"
	"sync"
)

func New() Handler {
	return Handler{
		Balancer: balancer.New(),
	}
}

type Handler struct {
	Balancer    *balancer.Balancer
	QueryStruct *queryStruct
	wg          sync.WaitGroup
}

type queryStruct struct {
	SqlQuery string `json:"sqlQuery"`
	Args     []any  `json:"args"`
}
