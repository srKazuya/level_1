package main

import "fmt"

func assertion(data []interface{}) {
	for _, s := range data {
		switch v := s.(type) {
		case int:
			fmt.Println("int:", v)
		case string:
			fmt.Println("string:", v)
		case bool:
			fmt.Println("bool:", v)
		case chan int:
			fmt.Println("channel:", v)
		}
	}
}

func main() {
	var ch chan int

	dataV := []interface{}{1, "apple", true, ch}
	assertion(dataV)
}
