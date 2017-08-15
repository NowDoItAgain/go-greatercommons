package main

import "fmt"

const (
	a = 2014 + iota
	b = 2014 + iota
	c = 2014 + iota
	d = 2014 + iota
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
