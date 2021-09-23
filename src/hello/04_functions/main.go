package main

import (
	"fmt"
	"math"
	"strconv"
)

func greeting(name string, name2 string, age int) string {
	return "Hello " + name + "I am " + strconv.Itoa(age) + " Years old of age and I love " + name2
}

func getSum(a int, b int, c int) int {
	return (a + b*c) / 2

}

func formlar(a float64, b float64, c float64) (float64, float64) {
	var X1 float64 = (-b + (math.Sqrt((b * b) - 4*a*c))) / 2 * a
	var X2 float64 = (-b - (math.Sqrt((b * b) - 4*a*c))) / 2 * a

	return X1, X2
}

func main() {
	fmt.Println(greeting("Kelechi", "Favour", 45))
	fmt.Println(getSum(10, 20, 30))
	fmt.Println(formlar(6.0, 15.3, 4.5))
}
