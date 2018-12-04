package main

import "fmt"

// START OMIT
func main() {
	fmt.Println("After this, panic will start")
	panic("Panic occoured!")
	fmt.Println("This line will not appear")
}

// END OMIT
