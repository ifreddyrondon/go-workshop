package main

import (
	"fmt"
	"time"
)

func main() {
	oreChan := make(chan int)
	minedOreChan := make(chan int)
	// START OMIT
	theMine := []string{"rock", "ore", "ore", "rock", "ore"}
	go func(mine []string) { // Finder
		for i, item := range mine {
			if item == "ore" {
				fmt.Printf("Found: ore %d\n", i)
				oreChan <- i //send
			}
		}
	}(theMine)
	go func() { // Ore Breaker
		for i := 0; i < 3; i++ {
			foundOre := <-oreChan //receive
			fmt.Printf("Miner: received ore %d. Mining...\n", foundOre)
			minedOreChan <- foundOre //send to minedOreChan
		}
	}()
	go func() { // Smelter
		for i := 0; i < 3; i++ {
			minedOre := <-minedOreChan
			fmt.Printf("Smelter: received mined ore %d. Smelting...\n", minedOre)
		}
	}()
	// END OMIT
	time.Sleep(time.Second * 5)
}
