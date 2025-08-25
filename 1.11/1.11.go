package main

import (
	"fmt"
)

func main() {
	in := make(map[int]struct{})
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{2, 3, 4, 1, 20, 6, 1}
	seen := make(map[int]struct{})

	res := make([]int, 0)
	for _, v := range a {
		in[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := in[v]; ok {
			if _, inside := seen[v]; !inside{
				res = append(res, v)
				seen[v] = struct{}{}
			} 
		}
	}

	fmt.Println(res)

}
