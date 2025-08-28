package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Первое число")
	var x big.Int
	fmt.Fscan(in, &x)

	fmt.Println("Операция")
	var op string
	fmt.Fscan(in, &op)

	fmt.Println("Второе число")
	var y big.Int
	fmt.Fscan(in, &y)

	answer := new(big.Int)
	switch op {
	case "+":
		answer.Add(&x, &y)
	case "-":
		answer.Sub(&x, &y)
	case "/":
		answer.Div(&x, &y)
	case "*":
		answer.Mul(&x, &y)
	default:
		fmt.Println("Неизвестная операция")
	}

	fmt.Println(answer)
}
