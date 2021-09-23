package main

import (
	"fmt"
)

func main() {
	X := 10
	y := 20

	if X >= y {
		fmt.Printf("%d is less than or equal to %d\n", X, y)
	} else {
		fmt.Printf("%d is less than %d", y, X)
	}

	// else if
	color := "red"

	if color == "red" {
		fmt.Println("The color is red")
	} else if color == "blue" {
		fmt.Println("The color is blue")
	} else {
		fmt.Println("Color is either blue or red")
	}

	// switch

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("Color is blue")

	default:
		fmt.Println("Color is either blue or red")
	}
}
