package main

import (
	"fmt"
	"time"
)

// START OMIT
func worker(done chan bool) {
	fmt.Println("working...")

	time.Sleep(time.Second)

	fmt.Println("worker done")

	done <- true
}

func main() {
	done := make(chan bool, 1)

	go worker(done)

	<-done
	fmt.Println("end")
}

// END OMIT
