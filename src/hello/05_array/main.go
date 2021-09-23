package main

import (
	"fmt"
)

func main() {
	// array
	var frutArr [2]string

	// Assign Values
	frutArr[0] = "Apple"
	frutArr[1] = "Orange"

	// declear and assign
	fruitSlice := []string{"Apple", "Orange", "Grape", "Water Mellon"}

	fmt.Println(frutArr[1])
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[0:2])
}
