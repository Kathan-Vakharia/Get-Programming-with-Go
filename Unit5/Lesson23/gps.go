package main

import (
	"fmt"
	"math"
)

//world with given volumetric radius
type world struct{ radius float64 }

//distance calculates distance between two locations
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))

	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

//rad function converts degree to radian
func rad(degree float64) float64 {
	return math.Pi * degree / 180

}

//location defines coordinates in degree
type location struct {
	name      string
	lat, long float64
}

//newLocation creates new location
func newLocation(name string, lat, long float64) location {
	return location{
		name: name,
		lat:  lat,
		long: long,
	}
}

//description to return information about location
func (l location) description() string {
	return fmt.Sprintf("%s, latitude: %.3f, longitude: %.3f", l.name, l.lat, l.long)
}

//gps for identifying the location
type gps struct {
	source, destination location
	world
}

//distance calculates b/w source and destination
func (g gps) distance() float64 {
	//forwarding to world's distance method
	return g.world.distance(g.source, g.destination)

}

//message returns kilometeres left for destination
func (g gps) message() string {
	return fmt.Sprintf("%.3f km to %s", g.distance(), g.destination.description())
}

//rover type composed of gps type
type rover struct {
	//embedding gps type to rover
	gps 
}

//defining world
var (
	mars = world{radius: 3389.5}
)

func main() {
	bradburry := newLocation("Bradburry", -4.5895, 137.4417)
	elysium := newLocation("Elysium Planitia", 4.5, 135.9)
	gps := gps{
		source:      bradburry,
		destination: elysium,
		world:       mars,
	}

	curiosity := rover{gps: gps}

	fmt.Println(curiosity.message())

}
