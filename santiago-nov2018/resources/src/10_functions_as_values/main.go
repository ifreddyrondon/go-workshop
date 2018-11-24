package main

import (
	"fmt"
)

// START OMIT
type x func(int)

func printInt(v int) {
	fmt.Println(v)
}

func main() {
	f := getFunction()
	f(1)
	//executeFunction(printInt, 4)
}

func getFunction() func(int) {
	f := func(value int) {
		fmt.Printf("Execute function with value: %d \n", value)
	}
	return f
}

func executeFunction(f x, value int) {
	f(value)
}

// END OMIT
