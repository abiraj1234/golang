package main

import (
	"fmt"
	"sync"
)

const turns = 100000

type counter struct {
	count int
	sync.Mutex
}

func main() {
	counter := counter{count: 0}
	done := make(chan struct{})
	go func() {
		for i := 0; i < turns; i++ {
			counter.Lock()
			counter.count++
			counter.Unlock()
		}
		done <- struct{}{}
	}()
	go func() {
		for i := 0; i < turns; i++ {
			counter.Lock()
			counter.count--
			counter.Unlock()

		}
		done <- struct{}{}
	}()

	<-done
	<-done
	fmt.Println(counter.count)
}
