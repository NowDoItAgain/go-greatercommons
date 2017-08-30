package main

import "fmt"

func main() {

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	toyota := truck{
		vehicle: vehicle{
			doors: 4,
			color: "maroon",
		},
		fourWheel: true,
	}

	honda := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "silver",
		},
		luxury: false,
	}

	fmt.Println(toyota)
	fmt.Println(honda)
	fmt.Println("")
	fmt.Println(toyota.color)
	fmt.Println(honda.doors)

}
