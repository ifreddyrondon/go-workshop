package main

import (
	"errors"
	"fmt"
)

// START OMIT
func HandlePayment() error {
	return errors.New("payment failed")
}

func main() {
	if err := HandlePayment(); err != nil {
		if err.Error() == "payment failed" {
			fmt.Println("send notification")
		}
	}
}

// END OMIT
