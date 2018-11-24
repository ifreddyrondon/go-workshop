package main

import (
	"fmt"
)

// START OMIT
func DoSomething(v interface{}) {
	fmt.Println(v)
}

func main() {
	DoSomething("String")
	DoSomething('C')
	DoSomething(1)
	DoSomething(false)
}

// END OMIT
