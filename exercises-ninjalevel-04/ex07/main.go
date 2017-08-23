package main

import "fmt"

func main() {

	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "MoneyPenny", "Helloooo, James"}
	fmt.Println(x)
	fmt.Println(y)

	md := [][]string{x, y}
	fmt.Println(md)

	for i, key := range md {
		fmt.Println("record: ", i)
		for j, val := range key {
			fmt.Printf("\t index position: %v \t value: %v \n", j, val)
		}
	}
}
