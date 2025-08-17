package main

import (
	"fmt"
	"sync"
	"time"
)

var workers int= 10

func main() {
	ch := make(chan int)

	wg := &sync.WaitGroup{}


	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for v := range ch {
				fmt.Println(v)
			}

		}()
	}

	num := 0
	for {
		num++
		ch <- num
		time.Sleep(2 * time.Second)
	}
}
