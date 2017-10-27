package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hi, my name is", p.first, p.last, "\nI am", p.age, "years old")
}

func main() {

	p1 := person{
		first: "John",
		last:  "Wayne",
		age:   37,
	}

	p1.speak()

}
