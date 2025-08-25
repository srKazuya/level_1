package main

import (
	"fmt"
	"math"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempM := make(map[int][]float64)

	for _, v := range temps {
		tens := (int(math.Floor(v))/10%10)

		tempM[tens*10] = append(tempM[tens*10], v)
	}

	fmt.Println(tempM)
}