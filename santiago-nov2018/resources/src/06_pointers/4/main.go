package main

import (
	"fmt"
)

// START OMIT
func main() {
	a := new(int)
	*a = 10

	fmt.Printf("Before method   -> address: %v, value: %d \n", a, *a)
	increment(a)
	fmt.Printf("After method    -> address: %v, value: %d \n", a, *a)
}

func increment(a *int) {
	fmt.Printf("func inc a -> address: %v, value: %d \n", a, *a)
	b := new(int)
	*b = *a
	*b++
	a = b
	fmt.Printf("func inc a -> address: %v, value: %d \n", a, *a)
	fmt.Printf("func inc b -> address: %v, value: %d \n", b, *b)
}

// END OMIT
