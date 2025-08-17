package main

import "fmt"

type human struct {
	name    string
	age     int
	address string
}

type action struct {
	human
	act string
}

func NewPerson(age int, name, address string) human {
	return human{
		name:    name,
		age:     age,
		address: address,
	}
}

func NewAction(human human, act string) action {
	return action{
		human: human,
		act:   act,
	}
}

func (a action) Print() {
	fmt.Printf("Человек: %v, Делает: %v", a.human.name, a.act)
}

func main() {

	ivan := NewPerson(25, "Ivan", "Moscow")
	ivanAction := NewAction(ivan, "study")


	ivanAction.Print()
}
