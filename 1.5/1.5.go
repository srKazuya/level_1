package main

import (
	"fmt"
	"time"
)

func worker(n int, timer time.Timer) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i < n; i++ {
			select {
			case ch <- i:
			case <-timer.C:
				return
			}
		}
	}()
	return ch
}

func main() {
	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()
	
	ch := worker(10, *timer)
	
	for {
		select {
		case v, ok := <-ch:
			time.Sleep(1 * time.Second)
			if !ok {
				return
			}
			fmt.Println(v)
		case <-timer.C:
			return
		}
	}
}
