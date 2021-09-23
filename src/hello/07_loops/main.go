package main

import (
	"fmt"
)

func main() {
	// Long method
	i := 1
	x := 10
	for i <= x {

		fmt.Println(i)
		i++
	}
	// short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}
}
