package main

import (
	"fmt"
	"strings"
)

func main() {
	st := "snow dog sun"
	s := strings.Fields(st)

	left := 0
	right := len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	fmt.Println(s)
}
