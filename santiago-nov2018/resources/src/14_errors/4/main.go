package main

import (
	"errors"
	"fmt"
)

var ErrFailedPayment = errors.New("payment failed")

// START OMIT
func HandlePayment() error {
	return ErrFailedPayment
}

func main() {
	if err := HandlePayment(); err != nil {
		if err == ErrFailedPayment {
			fmt.Println("send notification")
		}
	}
}

// END OMIT
