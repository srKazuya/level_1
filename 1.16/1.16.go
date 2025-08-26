package main

import "fmt"

func quick(a []int, frst, last int) {
	if frst >= last {
		return 
	}
	pos := part(a, frst, last)
	quick(a, frst, pos-1)
	quick(a, pos+1, last)
}

func part(s []int, frst, last int) int {

	pivot := s[last]
	wall := frst - 1

	for elemI := frst; elemI < last; elemI++ {
		if s[elemI] < pivot {
			wall++
			s[elemI], s[wall] = s[wall], s[elemI]
		}
	}
	s[last], s[wall+1] = s[wall+1], s[last]
	return wall + 1
}

func main() {
	nums := []int{4, 2, 5, 1, 9, 3}
	quick(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
