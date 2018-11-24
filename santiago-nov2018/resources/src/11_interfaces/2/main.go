package main

import (
	"fmt"
)

// Superhero interface. If you want to be a superhero you must have a super power.
type Superhero interface {
	Powers()
}

type Spiderman struct {
	Alias string
	Age   int
}

func (s Spiderman) Powers() {
	fmt.Printf("%s powers: %s", s.Alias, "superhuman strength, speed, spider-sense")
}

// START OMIT
type Flash struct {
	Alias string
	Age   int
}

func (s Flash) Powers() {
	fmt.Printf("%s powers: %s", s.Alias, "speed")
}

func main() {
	superheros := []Superhero{
		Spiderman{Alias: "Spiderman", Age: 28},
		Flash{Alias: "Flash", Age: 24},
	}

	for _, hero := range superheros {
		hero.Powers()
		fmt.Println()
	}

}

// END OMIT
