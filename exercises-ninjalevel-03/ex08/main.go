package main

import "fmt"

func main() {
	switch "hello" {
	case "not me":
		fmt.Println("not going to print")
	case "hello":
		fmt.Println("Hi, I will print")

	}
}
