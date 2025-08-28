package main

import "fmt"

func main() {
	s := []int{1, 5, 2, 10, 3}
	indx := 1

	if len(s)-1 < indx {
		return
	}
	copy(s[indx:], s[indx+1:])
	s = s[:len(s)-1]

	fmt.Println(s)

}
