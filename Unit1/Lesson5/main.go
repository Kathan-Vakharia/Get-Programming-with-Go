// Capstone : Ticket To Mars
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	distane       = 62100000  //km
	secondsPerDay = 24 * 3600 //s
)

func main() {
	//seeding random for random results
	rand.Seed(time.Now().UnixNano())
	//declaring necessary variables
	var (
		speed     int
		price     int
		days      int
		tripType  string
		spaceline string
	)

	fmt.Printf("%-20s\tDays\tTripType\tPrice\n", "Spaceline")
	fmt.Println("=====================================================")

	for i := 0; i < 10; i++ {

		switch rand.Intn(3) {
		case 0:
			spaceline = "SpaceX"
		case 1:
			spaceline = "Virgin Galactic"
		case 2:
			spaceline = "Space Adventures"
		}

		//calculating speed b/w 16-30 km/s
		speed = 16 + rand.Intn(15)
		//calculating Days
		days = distane / speed / secondsPerDay
		//calculating price b/w 36 to 50 Million
		price = 20 + speed

		//Taking account of Type of trip
		if rand.Intn(2) == 0 {
			tripType = "One-way"
		} else {
			tripType = "Round-Trip"
			price *= 2
		}

		fmt.Printf("%-20s\t%-4v\t%-9v\t$ %3v\n", spaceline, days, tripType, price)

	}
}
