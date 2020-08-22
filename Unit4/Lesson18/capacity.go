package main

import "fmt"

const maxIterations = 1000
func main() {
	s := make([]int, 0)
	lastCap := cap(s)
	fmt.Println("Starting Capacity:", lastCap)

	for iteration:=0;iteration<maxIterations;iteration++{
		s  = append(s, iteration)
		if(cap(s) != lastCap){
			lastCap = cap(s)
			fmt.Printf("Capacity %4d at length %3d\n",lastCap,iteration+1)
		}
	}

}
