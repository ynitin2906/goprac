package oops

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func RunMutex() {
	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter Value:", counter.GetValue())
}
