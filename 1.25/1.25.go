package main

import (
	"fmt"
	"time"
)

func Sleep(t time.Duration) {
	ch := make(chan int)

	timer := time.NewTimer(t * time.Second)
	go func() {
		i := 0
		for {
			i++
			select {
			case ch <- i:
				fmt.Println(i)
			case <-timer.C:
				close(ch)
				return
			}
		}
	}()
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(v)
		case <-timer.C:
			fmt.Println("exited by timer")
			return
		}
	}

}

func main() {
	Sleep(5)
}
