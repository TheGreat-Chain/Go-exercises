/**
* Reproduce an Sqrt function, and compare it to the math.Sqrt() function
*/

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1.0
	for i := 0 ; i < 100 ; i++ {
		if z*z == x {
			break
		}
		z -= (z*z - x) / (2*z)	
	}
	return z
	
}

func main() {
	var x float64 = 98766
	fmt.Println(Sqrt(x))
	fmt.Println(math.Sqrt(x))
}
