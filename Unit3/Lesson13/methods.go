package main

import "fmt"

type fahrenheit float64
type celsius float64
type kelvin float64

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (c celsius) fahrenheit() fahrenheit {
	f := 9*c/5 + 32
	return fahrenheit(f)
}

func (f fahrenheit) celsius() celsius {
	c := (f - 32) * 5 / 9
	return celsius(c)
}
func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func main() {
	var k kelvin = 294.14
	fmt.Printf("%.2f° K is %.2f° C\n", k, k.celsius())
	fmt.Printf("%.2f° K is %.2f° F", k, k.fahrenheit())
}
