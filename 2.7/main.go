package main

import (
	"fmt"
	"math"
)

var k, p, v float64 = 32.5, 23.5, 16.9

func main() {
	fmt.Print(T())
}

func M() float64 {
	return p * v
}

func W() float64 {
	return math.Sqrt(k / M())
}

func T() float64 {
	return 6 / W()
}
