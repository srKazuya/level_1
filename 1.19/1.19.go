package main

import "fmt"

func main() {
	s := "главрыба"
	r := []rune(s)
	out := make([]rune, len(r))

	for i := range r{
		out[len(r)-1-i] = r[i]
	}

	fmt.Println(string(out))

}
