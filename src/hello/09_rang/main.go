package main

import (
	"fmt"
)

func main() {

	ids := []int{10, 11, 20, 39, 44, 54, 56, 67, 89, 78, 68}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
}
