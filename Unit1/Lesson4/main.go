//Go program to generate a Random Date
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {
	//seeing random for randomness
	rand.Seed(time.Now().UnixNano())
	var year, daysInMonth, month, day int
	fmt.Println("Era year month day")

	for i := 0; i < 10; i++ {
		//generating a random year of 21st Century
		year = 2000 + rand.Intn(101)
		month = rand.Intn(12) + 1

		switch month {
		case 2: //Feb
			if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}
		case 4, 6, 9, 11: //Apr Jun Sept Nov
			daysInMonth = 30
		default: // Jan Mar May Jul Aug Oct Dec
			daysInMonth = 31
		}
		day = rand.Intn(daysInMonth) + 1
		fmt.Printf("%-3v %4v %5v %02v\n",era,year, month,day)

	}
}
