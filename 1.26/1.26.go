package main

import (
	"fmt"
	"unicode"
)

func isOk(s string) bool {
	m := make(map[rune]struct{})

	for _, v := range s {
		if _, ok := m[unicode.ToLower(v)]; !ok {
			m[v] = struct{}{}
		} else {
			return false
		}
	}

	return true
}

func main() {
	f := isOk("abfDe")
	t := isOk("abfDAe")

	fmt.Println(t)
	fmt.Println(f)
}
