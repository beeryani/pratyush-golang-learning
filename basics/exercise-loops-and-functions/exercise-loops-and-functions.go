package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	count := 0
	z := 1.0
	for count < 10{
		z -= (z*z - x) / (2*z)
		fmt.Printf("-- %g\n",z)
		count += 1
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
