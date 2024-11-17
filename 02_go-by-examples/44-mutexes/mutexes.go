package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		counters: map[string]int{"a":0, "b":0},
	}

	var wg sync.WaitGroup

	doIncreament := func(name string, n int){
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()

	}

	wg.Add(3)
	go doIncreament("a", 10000)
	go doIncreament("a", 10000)
	go doIncreament("b", 10000)

	wg.Wait()

	fmt.Println(c.counters)
}