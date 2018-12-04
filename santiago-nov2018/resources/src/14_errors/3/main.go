package main

import (
	"fmt"
)

// START OMIT
type ErrFailedPayment struct {
	Msg, Amount string
}

func (e *ErrFailedPayment) Error() string { return e.Msg }
func HandlePayment() error {
	return &ErrFailedPayment{Msg: "payment failed", Amount: "100000"}
}

func main() {
	err := HandlePayment()
	if err != nil {
		if err, ok := err.(*ErrFailedPayment); ok {
			fmt.Printf("payment failed. Amount: %v", err.Amount)
		}
	}
}

// END OMIT
