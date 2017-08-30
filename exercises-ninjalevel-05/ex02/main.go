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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.favIce {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}

}
