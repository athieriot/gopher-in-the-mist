package main

import (
	"fmt"
	"math"
)

// Here is a function with parameters and return type // HL
func pow(x, n, lim float64) float64 {
	// Ifs can have an init part (optional) // HL
	// Oh and no brackets here // HL
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}