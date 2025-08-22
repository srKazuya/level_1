package main

import (
	"fmt"
	"sync"
	"time"
)

var arr = []int{1, 2, 3, 4, 5}

func writer(a []int) <-chan int {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	for _, v := range a {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			ch <- val
		}(v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}

func multiply(in <-chan int) <-chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
				defer wg.Done()
			for v := range in {
				out <- (v * 2)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func reader(in <-chan int) []int {
	var arr []int

	for v := range in {
		arr = append(arr, v)
	}
	
	return arr
}

func main() {

	res := reader(multiply(writer(arr)))
	fmt.Println(res)
}
