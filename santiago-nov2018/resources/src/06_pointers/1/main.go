package main

import (
	"fmt"
)

// START OMIT
func main() {
	a := 10

	fmt.Printf("Before method    -> value: %d \n", a)
	increment(a)
	fmt.Printf("After  method    -> value: %d \n", a)
}

func increment(a int) {
	a++
	fmt.Printf("Increment method -> value: %d \n", a)
}

// END OMIT
