package main

import (
	"fmt"
)

// START SUPERHERO OMIT

// Superhero interface. If you want to be a superhero you must have a super power.
type Superhero interface {
	Powers()
}

// END SUPERHERO OMIT

type Spiderman struct {
	Alias string
	Age   int
}

func (s Spiderman) Powers() {
	fmt.Printf("%s powers: %s", s.Alias, "superhuman strength, speed, spider-sense")
}

func main() {
	var spiderman Superhero = &Spiderman{Alias: "Spiderman", Age: 28}
	spiderman.Powers()
}

// END OMIT
