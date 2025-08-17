package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var workers int = 10

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()

	ch := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case v, ok := <-ch:
					if !ok {
						return
					}
					fmt.Println(v)
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	num := 0
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("program exited by context")
			return
		case <-time.After(2 * time.Second):
			num++
			ch <- num
		}
	}

	fmt.Println("program ended")
	close(ch)
	wg.Wait()
}
