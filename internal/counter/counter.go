package counter

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func New() *Counter {
	return &Counter{}
}

type Counter struct {
	MainCounter   atomic.Int64
	Slave1Counter atomic.Int64
	Slave2Counter atomic.Int64
}

func (c *Counter) IncrementCounter(ops *atomic.Int64) {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		ops.Add(1)
		wg.Done()
	}()

	wg.Wait()

}

func (c *Counter) DecrementCounter(ops *atomic.Int64) {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		ops.Add(-1)
		wg.Done()
	}()

	wg.Wait()

}

func (c *Counter) LoadCounter() (m, s1, s2 int64) {

	m = c.MainCounter.Load()
	s1 = c.Slave1Counter.Load()
	s2 = c.Slave2Counter.Load()

	return m, s1, s2
}

func (c *Counter) ConvertOnString(u int64) string {
	return fmt.Sprintf("%d", u)
}
