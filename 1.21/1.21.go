package main

import (
	"fmt"
)

type Family interface {
	Speaking() string
}

type Son struct {
	Name string
}

func (s *Son) Speaking() string {
	return fmt.Sprintf("%s разговаривает", s.Name)
}

type Dog struct {
	Name string
}

func (d *Dog) Barking() string {
		return fmt.Sprintf("%s гавкает", d.Name)
}

type DogAdapter struct {
	dog *Dog
}

func (a *DogAdapter) Speaking() string {
	return a.dog.Barking()
}

func main() {
	family := []Family{
		&Son{Name: "Иван"},
		&DogAdapter{dog: &Dog{Name: "Крапинка"}},
	}

	for _, v := range family{
		fmt.Println(v.Speaking())
	}
}
