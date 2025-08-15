package main

import (
	"fmt"
	"math"
	"sync"
)

var arr [5]int = [5]int{2, 4, 6, 8, 10}

func main() {

	wg := &sync.WaitGroup{}

	for _, v := range arr {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			res := math.Pow(float64(n), 2)
			fmt.Println(res)
		}(v)
	}
	wg.Wait()
}
