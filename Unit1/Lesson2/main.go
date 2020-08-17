package main

import (
	"fmt"
)

func main() {
	//* Listing 2.1 Check weight and age on Mars
	// fmt.Println("My weight on Mars: ", 55*0.3783, " (kgs)")
	// fmt.Println("My age on Mars: ", 19*365.2425/687, " years")

	//* Format verb %v
	//p -ve values adds padding to right
	//p +ve values add padding to the left
	// fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	// fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

	//* QC 2.3 Time for SpaceX to reach Mars
	// const hoursInADay = 24 //hours
	// var (
	// 	speed = 100800   //km/hr
	// 	dist  = 96300000 // kms
	// )
	// fmt.Printf("Number of days taken by SpaceX to reach Mars are %v", (dist/speed)/hoursInADay)
	//* Declaring multple variables in same line
	// var num1, num2 = 5, 8
	// println(num1, num2)

	//* rand package
	// rand.Seed(time.Now().UnixNano())
	// var num = rand.Intn(10) + 1
	// fmt.Println("The random int generated is ", num)

	//* QC 2.6
	//- to change seed dynamically
	// rand.Seed(time.Now().UnixNano())
	// var dist = 34500001 + rand.Intn(401000000)
	// println("Random Distance: ", dist)

	//* Experiment: Malacandra.go
	const (
		hoursInDay = 24
		days       = 28
	)
	var dist = 56000000 //km

	//- declaring using implicit initialization syntax
	speed := dist / (days * hoursInDay)
	fmt.Printf("The speed rquired is %v km/hr", speed)

}
