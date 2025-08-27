package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	num int
}

func main() {
	c := new(Counter)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			c.num++
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(c.num)
}
