//Piggy bank with whole numbers
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	piggybank := 0

	//2000 cents -> 20$
	for piggybank < 2000 {
		switch rand.Intn(3) {
		case 0: //nickels
			piggybank += 5
		case 1: //dimes
			piggybank += 10
		case 2: //quarters
			piggybank += 25
		}
		dollar := piggybank / 100
		cents := piggybank % 100
		fmt.Printf("Piggy Count: $ %02d.%02d\n", dollar, cents)

	}
}
