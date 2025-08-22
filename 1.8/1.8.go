package main

import (
	"bufio"
	"fmt"
	"os"
)

var num int64

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Println("Введите число")
	fmt.Fscan(in, &num)

	var i int
	fmt.Println("Введите номер бита")
	fmt.Fscan(in, &i)

	var t int
	fmt.Println("0 1")
	fmt.Fscan(in, &t)
	
	switch t {
	case 0:
		nul := num &^ (1 << i)
		fmt.Printf("Новое число: %d\n", nul)
		fmt.Fprintln(out, nul)
	case 1:
		one := num | (1 << i)
		fmt.Printf("Новое число: %d\n", one)
	}
}
