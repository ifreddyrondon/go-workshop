package main

import "fmt"

// START OMIT
func main() {
	fmt.Println("Counting")
	var i int
	for i <= 5 {
		i++
		defer fmt.Println(i)
	}
	i++
	defer fmt.Printf("i: %d\n", i)
}

// END OMIT
