package main

import "fmt"

func bin(s []int, t, offset int) int {
	if len(s) == 0 {
		return -1
	}

	mid := len(s) / 2
	left := s[:mid]
	right := s[mid+1:]
	
	if s[mid] > t {
		return bin(left, t, offset)
	}
	if s[mid] < t {
		return bin(right, t, offset+mid+1)
	}

	return offset + mid
}

func main() {
	nums := []int{1, 2}
	target := 2

	res := bin(nums, target, 0)
	fmt.Println(res)
}
