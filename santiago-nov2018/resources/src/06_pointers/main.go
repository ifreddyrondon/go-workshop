package main

import (
	"fmt"
)

func main() {
	// START OMIT
	x := 4
	fmt.Printf("X Value = %v \n", x)

	y := new(int)
	fmt.Printf("Y = %v \n", y)
	fmt.Printf("Y Value = %v \n", *y)

	*y = 4
	fmt.Printf("Y Value = %v \n", *y)

	if y == &x {
		fmt.Println("X & Y are equals")
	} else {
		fmt.Println("X & Y are not equals")
	}
	// END OMIT
}
