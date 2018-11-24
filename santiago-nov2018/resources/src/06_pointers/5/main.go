package main

import (
	"fmt"
)

// START OMIT
func main() {
	a := new(int)
	*a = 10
	fmt.Printf("Before method   -> p value: %v, p address: %v, value: %d \n", a, &a, *a)
	increment(a)
	fmt.Printf("After method    -> p value: %v, p address: %v, value: %d \n", a, &a, *a)
}

func increment(a *int) {
	fmt.Printf("func inc  a -> p value: %v, p address: %v, value: %d \n", a, &a, *a)
	b := new(int)
	*b = *a
	*b++
	a = b
	fmt.Printf("func inc  a -> p value: %v, p address: %v, value: %d \n", a, &a, *a)
	fmt.Printf("func inc  b -> p value: %v, p address: %v, value: %d \n", b, &b, *b)
}

// END OMIT
