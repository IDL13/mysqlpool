package handler

import (
	"mysqlpool/pkg/mysqlconn"
	"sync"
)

func New() Handler {
	return Handler{
		Compound: mysqlconn.New(),
	}
}

type Handler struct {
	Compound    *mysqlconn.Compound
	QueryStruct *queryStruct
	wg          sync.WaitGroup
}

type queryStruct struct {
	SqlQuery string `json:"sqlQuery"`
	Args     []any  `json:"args"`
}
