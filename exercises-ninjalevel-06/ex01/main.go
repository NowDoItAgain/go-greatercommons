package main

import (
	"fmt"
)

func foo() int {
	return 33
}

func bar() (int, string) {
	return 5, "little monkeys"
}

func main() {

	fmt.Println(foo())
	fmt.Println(bar())

}
