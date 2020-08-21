package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var offset kelvin = 5

	println()
	//This offset is passed by value
	sensor := calibrate(realSensor, offset)
	fmt.Println("Real Sensor value before changing offset", sensor())

	//changing offset
	offset ++
	fmt.Println("Real Sensor value after changing offset", sensor())

	sensor = calibrate(fakeSensor, offset)
	for i:=0; i<10;i++{
		fmt.Println("Fake sensor value:", sensor())
	}

}
