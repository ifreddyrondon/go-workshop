package main

import (
	"fmt"
	"os"
)

// START OMIT
func main() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

// END OMIT
