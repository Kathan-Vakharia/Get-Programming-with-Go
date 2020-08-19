//* Distance from canis.go

package main

import "fmt"

const (
	dist          = 236000000000000000 //km
	secondsPerDay = 86400              //seconds
	lightSpeed    = 299792             //km/s
	daysInYear    = 365                //days

)

func main() {
	numberOfLightYears := dist / lightSpeed / secondsPerDay / daysInYear

	fmt.Println("The canis Dwarf Galaxy is", numberOfLightYears, "light years away")

}
