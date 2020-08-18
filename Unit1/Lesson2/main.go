package main

import (
	"fmt"
)

func main() {
	//* Experiment: Malacandra.go
	const (
		hoursInDay = 24
		days       = 28
	)
	var dist = 56000000 //km

	//- declaring using implicit initialization syntax
	speed := dist / (days * hoursInDay)
	fmt.Printf("The speed required is %v km/hr", speed)

}
