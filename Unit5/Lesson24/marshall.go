package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//coordinate represents a latitude or longitude in dmsh
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

//MrshalJSON to satisfy json.Marshaler interface
func (c coordinate) MarshalJSON() ([]byte, error) {

	return json.MarshalIndent(map[string]interface{}{
		"decimal":    c.decimal(),
		"dmsh":       fmt.Sprintf(`%.0f°%.0f'%0.0f" %c`, c.d, c.m, c.s, c.h),
		"degrees":    c.d,
		"minutes":    c.m,
		"seconds":    c.s,
		"hemisphere": c.h,
	}, "", " ")
}

//location defines a location in latitude and longitude
type location struct {
	name      string
	lat, long coordinate
}

//newLocation to create a new location
func newLocation(name string, lat, long coordinate) location {
	return location{
		name: name,
		lat:  lat,
		long: long,
	}
}

//String method to satisty Stringer Interface 
func (l location) String() string {
	lat, e1 := l.lat.MarshalJSON()
	long, e2 := l.long.MarshalJSON()

	if e1 != nil || e2 != nil {
		fmt.Println("Error parsing json")
		os.Exit(1)
	}
	return fmt.Sprintf("%s\n%s\n%s", l.name, string(lat), string(long))
}


func main() {
	location := newLocation("Elysium", coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0.0, 'E'})

	fmt.Println(location)
}
