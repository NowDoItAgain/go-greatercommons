package main

import "fmt"

func main() {

	watch := struct {
		hands   bool
		digital bool
		timer   bool
		windup  string
	}{
		hands:   true,
		digital: false,
		timer:   false,
		windup:  "yes, I can be wound",
	}

	fmt.Println(watch)
	fmt.Println(watch.digital)

}
