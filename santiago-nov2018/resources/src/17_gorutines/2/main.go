package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello gouroutines!")
}

// START OMIT
func main() {
	go sayHello()
	// Do other things
}

// END OMIT
