package main

import (
	"fmt"
)

// START OMIT
type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func main() {
	var err error = &CustomError{Message: "It is a custom error"}
	fmt.Printf("Error: %s \n", err.Error())
}

// END OMIT
