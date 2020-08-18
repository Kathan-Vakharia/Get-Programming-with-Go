//* Static Guessing Game
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var userGuess, _ = strconv.Atoi(os.Args[1])

func main() {
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("##############################")
	var steps = 0

	//seeding random for better results
	rand.Seed(time.Now().UnixNano())

	//- Execute till generated random number is not equal to guess
	for {
		computerGuess := rand.Intn(100) + 1
		if computerGuess < userGuess {
			fmt.Println("Computers Guess", computerGuess)
			fmt.Println("Too Low!")
		} else if computerGuess > userGuess {
			fmt.Println("Computers Guess", computerGuess)
			fmt.Println("Too High!")
		} else if computerGuess == userGuess {
			//When guess matches
			fmt.Println("Ahan!")
			fmt.Printf("Computer's Guess: %v ,\nMyGuess: %v\n", computerGuess, userGuess)
			break
		}
		time.Sleep(time.Millisecond * 500)
		steps++
	}
	fmt.Printf("The computer guessed in %v tries", steps)

}
