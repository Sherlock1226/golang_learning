package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	sum := float64(x)

	if v := math.Sqrt(sum); v < 100 {
		if sum == v*v {
			return v
		}

	}

	return 0
}
func main() {

	fmt.Println(Sqrt(4))  // 2
	fmt.Println(Sqrt(9))  // 3
	fmt.Println(Sqrt(16)) // 4
	fmt.Println(Sqrt(10)) // 0
	fmt.Println(Sqrt(25)) // 0
	fmt.Println(Sqrt(36)) // 0
	fmt.Println(Sqrt(24)) // 0

}
