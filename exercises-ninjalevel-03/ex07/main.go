package main

import "fmt"

func main() {
	car := 15
	v := "This is car number "
	if car != 15 {
		fmt.Println(v, car)
	} else if car == 15 {
		fmt.Println("Car number 15 is looking good!")
	} else {
		fmt.Println("This is not a car.")
	}
}
