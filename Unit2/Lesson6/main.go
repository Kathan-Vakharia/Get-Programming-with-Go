package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	//starting from zero
	piggybank := 0.0

	for piggybank < 20 {

		switch rand.Intn(3) {
		case 0:
			piggybank += 0.05
		case 1:
			piggybank += 0.10
		case 2:
			piggybank += 0.25
		}
		fmt.Printf("Running Balance of Piggy %.2f\n", piggybank)
	}
}
