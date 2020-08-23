package main

import "fmt"

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

type location struct {
	vehicleName, landingSite string
	lat, long                float64
}

//newLocation returns a location in decimal degrees format
func newLocation(vehicleName, landingSite string, latDMS, longDMS coordinate) location {
	return location{
		vehicleName: vehicleName,
		landingSite: landingSite,
		lat:         latDMS.decimal(),
		long:        longDMS.decimal(),
	}
}

func main() {
	locations := []location{
		newLocation("Spirit", "Columbia Memorial Station", coordinate{14, 34, 6.2, 's'}, coordinate{175, 28, 21.5, 'e'}),
		newLocation("Opportunity", "Challenger Memorial Station", coordinate{1, 56, 46.3, 's'}, coordinate{354, 28, 24.2, 'e'}),
		newLocation("Curiosity", "Bradburry Landin", coordinate{4, 35, 22.2, 's'}, coordinate{137, 26, 30.1, 'e'}),
		newLocation("Insight", "Elysium Planitia", coordinate{4, 30, 0.0, 'n'}, coordinate{135, 54, 0, 'e'}),
	}
	fmt.Printf("%s\t%-24s\t%s\t%s\n", "Rover/Vehichle", "Landing site", "Latitude", "Longitude")
	for _, val := range locations {
		fmt.Printf("%-14s  %-30s  %-15.2f  %.2f\n", val.vehicleName, val.landingSite, val.lat, val.long)
	}
}
