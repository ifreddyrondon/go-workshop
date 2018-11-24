package main

import (
	"fmt"
)

// Superhero interface. If you want to be a superhero you must have a super power.
type Superhero interface {
	Powers()
}

// START OMIT
type Superman struct {
	Alias string
	Age   int
}

func (s *Superman) Powers() {
	fmt.Printf("%s powers: %s", s.Alias, "invulnerability, heat vision, flight, speed")
}

type SuperBoy struct {
	Superman
}

func main() {
	superboy := &SuperBoy{}
	superboy.Alias = "Superboy"
	superboy.Age = 24

	superboy.Powers()
}

// END OMIT
