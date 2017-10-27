package main

import (
	"fmt"
)

func foo(xi ...int) int {

	total := 0
	for _, v := range xi {

		total += v
	}
	return total
}

func bar(xx []int) int {
	sum := 0
	for _, v := range xx {
		sum += v
	}
	return sum
}

func main() {

	y := []int{5, 4, 3, 2, 1}
	fmt.Println(foo(y...))

	ix := []int{6, 7, 8, 9, 10}
	fmt.Println(bar(ix))

}
