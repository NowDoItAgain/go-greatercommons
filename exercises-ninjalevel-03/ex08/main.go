package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("not going to print")
	case true:
		fmt.Println("Hi, I will print")

	}
}
