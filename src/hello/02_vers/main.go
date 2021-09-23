package main

import "fmt"

func main() {
	var name = "Damian"
	var age = 40
	isCool := true
	// Shorthand variables
	mobile, email, address := "08133287426", "damiantopaz@gmail.com", "Ikeja lagos"

	fmt.Println(name, age)
	fmt.Printf("%T\n", isCool)
	fmt.Println(mobile, email, address)
}
