package main

import "fmt"

func main() {
	const (
		a int     = 33
		b float64 = 9.143
		c         = "Hello"
		d         = false
	)

	fmt.Println(a, b, c, d)
	fmt.Printf("%T\t%T\t%T\t%T\t", a, b, c, d)
}
