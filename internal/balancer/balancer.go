package balancer

import (
	"mysqlpool/pkg/mysqlconn"
)

func New() *Balancer {
	return &Balancer{
		Connections: mysqlconn.New(),
	}
}

type Balancer struct {
	Connections *mysqlconn.Compound
}

func (b *Balancer) GetHealth() map[string]bool {
	status := make(map[string]bool)

	status["main"] = b.Connections.Main.HealthinessProbe
	status["slave1"] = b.Connections.Slave1.HealthinessProbe
	status["slave2"] = b.Connections.Slave2.HealthinessProbe

	return status
}
