package main

import (
	"fmt"
	"math"
)

//degToRad converts degree value to radian value
func degToRad(deg float64) float64 {
	return math.Pi * deg / 180

}

type coordinate struct {
	d, m, s float64
	h       rune
}

//decimal method converts from dmsh -> decimal degree
func (c coordinate) decimal() float64 {
	sign := 1.0

	//changing sign depending on hemisphere
	switch c.h {
	case 'S', 's', 'W', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

//location represents location in form of latitude,longitude
type location struct {
	lat, long float64
}

//newLocation created a location with give coordinates
func newLocation(lat, long coordinate) location {
	return location{
		lat:  lat.decimal(),
		long: long.decimal(),
	}
}

//world type for new world
type world struct {
	radius float64
}

//distance method to calculate distance b/w two places on world
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(degToRad(p1.lat))
	s2, c2 := math.Sincos(degToRad(p2.lat))
	clong := math.Cos(degToRad(p1.long - p2.long))

	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

//defining worlds
var (
	mars  = world{radius: 3389.5}
	earth = world{radius: 6371}
)

func main() {

	spirit := newLocation(coordinate{14, 34, 6.2, 's'}, coordinate{175, 28, 21.5, 'e'})
	opportunity := newLocation(coordinate{1, 56, 46.3, 's'}, coordinate{354, 28, 24.2, 'e'})
	curiosity := newLocation(coordinate{4, 35, 22.2, 's'}, coordinate{137, 26, 30.1, 'e'})
	insight := newLocation(coordinate{4, 30, 0.0, 'n'}, coordinate{135, 54, 0, 'e'})

	fmt.Printf("Distance between Spirit and Opportunity %5.2f km\n", mars.distance(spirit, opportunity))
	fmt.Printf("Distance between Spirit and Curiosity %5.2f km\n", mars.distance(spirit, curiosity))
	fmt.Printf("Distance between Spirit and Curiosity %5.2f km\n", mars.distance(spirit, insight))
	fmt.Printf("Distance between Opportunity and Curiosity %5.2f km\n", mars.distance(opportunity, curiosity))
	fmt.Printf("Distance between Opportunity and Insight %5.2f km\n", mars.distance(opportunity, insight))
	fmt.Printf("Distance between Curiosity and Insight %5.2f km\n", mars.distance(curiosity, insight))

	println("-----------------------------------------------------------------")
	
	london := newLocation(coordinate{51, 30, 0, 'n'}, coordinate{0, 8, 0, 'w'})
	paris := newLocation(coordinate{48, 52, 0, 'n'}, coordinate{2, 21, 0, 'e'})
	fmt.Printf("Distance between London and Paris %5.2f km\n", earth.distance(london, paris))

	println("-----------------------------------------------------------------")
	
	mountSharp := newLocation(coordinate{5, 4, 48, 's'}, coordinate{137, 51, 0, 'e'})
	olympusMons := newLocation(coordinate{18, 39, 0, 'n'}, coordinate{226, 12, 0, 'e'})
	fmt.Printf("Distance between Mount Sharp and Olympus Mons %5.2f km\n", mars.distance(mountSharp, olympusMons))

	println("-----------------------------------------------------------------")

	ahmedabad := newLocation(coordinate{23, 2, 1.9068, 'n'}, coordinate{72, 35, 6.0792, 'e'})
	delhi := newLocation(coordinate{28, 38, 41.28, 'n'}, coordinate{77, 13, 0.1956, 'e'})
	fmt.Printf("Distance between Ahmedabad and Delhi %5.2f km\n", earth.distance(ahmedabad, delhi))

}
