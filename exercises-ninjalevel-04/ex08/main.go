package main

import "fmt"

func main() {

	m := map[string][]string{
		"bond_james": []string{"shaken, not stirred", "martinis", "women"}, "moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr": []string{"being evil", "ice cream", "sunsets"},
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
