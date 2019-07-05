package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sqrt(2), math.Sqrt(2))
	fmt.Println(sqrt(2212), math.Sqrt(2212))
}

func sqrt(x float64) float64 {
	z := float64(1)
	var adj float64
	for i := 0; i < 100000; i++ {
		adj = (z*z - x) / (2 * z)
		if x == 0 || math.IsNaN(x) || math.IsInf(x, 0) {
			return x
		} else if math.Abs(z*z-x) <= 0.0001 {
			fmt.Printf("took %d iterations to get the result\n", i)
			return z
		} else if math.IsInf(adj, 0) {
			return z
		} else {
			z -= adj
		}
	}
	return z
}
