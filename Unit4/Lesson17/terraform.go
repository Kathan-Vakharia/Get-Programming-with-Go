package main

import "fmt"

//Planets is used to hold string slice of planets
type Planets []string

//terraform method of Planets type manipulates the slice of planets
func (p Planets) terraform() {
	for index := range p {
		p[index] = "New " + p[index]
	}
}
func main() {

	planets := Planets{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	planets.terraform()
	//%q -> adds quotes
	fmt.Printf("%q",planets)
}
