package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello gouroutines!")
}

// START OMIT
func main() {
	go sayHello()
	// Do other things
	time.Sleep(time.Second)
}

// END OMIT
