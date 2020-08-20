package main

import "fmt"



func main() {
	// "\b" -> backspace escape sequence
	fmt.Println("233° K is", kelvinToCelsius(233),"\b° C")
	fmt.Println("0° K is",kelvinToFahrenheit(0),"\b° F")
}

//kelvinToCelcius converts Kelvin to Celsius
func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

//celsiusToFahrenheit converts Celsius to Fahrenheit
func celsiusToFahrenheit(c float64) float64 {
	return (9.0*c)/5.0 + 32
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}
