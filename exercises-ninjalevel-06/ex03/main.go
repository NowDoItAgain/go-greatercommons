package main

import (
	"fmt"
)

func first() {
	fmt.Println("I'm first in line, but I will print last because I'm Deferred!")
}

func last() {
	fmt.Println("I will print first even though I'm last in line")
}

func main() {
	defer first()
	last()
}
