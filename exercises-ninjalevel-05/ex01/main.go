package main

import "fmt"

func main() {

	type person struct {
		firstName string
		lastName  string
		favIce    []string
	}

	p1 := person{
		firstName: "Elmer",
		lastName:  "Fudd",
		favIce: []string{
			"vanilla",
			"chocolate",
		},
	}

	p2 := person{
		firstName: "Betty",
		lastName:  "Boop",
		favIce: []string{
			"strawberry",
			"peach",
		},
	}

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)

	for k, v := range p1.favIce {
		fmt.Println(k, v)
	}

	fmt.Print("\n=======\n\n")

	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)

	for k, v := range p2.favIce {
		fmt.Println(k, v)
	}

}
