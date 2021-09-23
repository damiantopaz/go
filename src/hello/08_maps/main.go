package main

import (
	"fmt"
)

func main() {
	emails := map[string]string{
		"damian": "damiantopaz@gmail.com",
		"mike":   "mike@gmail.com",
		"emma":   "emma@gmail.com",
		"uche":   "urchymanny@gmail.com",
	}
	fmt.Println(emails["damian"])

	// adding to map
	emails["chima"] = "chima@gmail.com"

	fmt.Println(emails)

}
