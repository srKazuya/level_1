package main

import "fmt"

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})
	res := make([]string, 0, len(s))

	for _, w := range s {
		if _, ok := set[w]; !ok {
			set[w] = struct{}{}
			res = append(res, w)
		}
	}
	fmt.Println(res)

}
