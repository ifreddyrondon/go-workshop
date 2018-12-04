package main

import (
	"fmt"
)

// START OMIT
var theMine = []string{"rock", "ore", "ore", "rock", "ore"}

func finder(mine []string) []string {
	foundOre := []string{}
	for _, v := range mine {
		fmt.Printf("from mine: %v\n", v)
		if v == "ore" {
			foundOre = append(foundOre, v)
		}
	}
	return foundOre
}
func miner(foundOre []string) []string {
	minedOre := []string{}
	for i := range foundOre {
		minedOre = append(minedOre, fmt.Sprintf("minedOre%d", i))
	}
	return minedOre
}
func main() {
	foundOre := finder(theMine)
	minedOre := miner(foundOre)
	fmt.Println(minedOre)
}

// END OMIT
