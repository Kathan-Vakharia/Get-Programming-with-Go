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

	return json.Marshal(map[string]interface{}{
		"decimal":    c.decimal(),
		"dmsh":       fmt.Sprintf(`%.0f°%.0f'%0.0f" %c`, c.d, c.m, c.s, c.h),
		"degrees":    c.d,
		"minutes":    c.m,
		"seconds":    c.s,
		"hemisphere": c.h,
	})
}

//location defines a location in latitude and longitude
type location struct {
	Name string     `json:"name"`
	Lat  coordinate `json:"lat"`
	Long coordinate `json:"long"`
}

//newLocation to create a new location
func newLocation(name string, lat, long coordinate) location {
	return location{
		Name: name,
		Lat:  lat,
		Long: long,
	}
}

func main() {
	location := newLocation("Elysium", coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0.0, 'E'})

	bytes, err := json.MarshalIndent(location,""," ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(string(bytes))
	}
}
