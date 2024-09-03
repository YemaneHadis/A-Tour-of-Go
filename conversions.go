package main

import (
	"fmt"
	"math"
)

func main() {

	var x, y int = 2, 32

	var f float64 = math.Sqrt(float64(x*x + y*y))
	var u uint = uint(f)
	fmt.Println(x, u, f)
	// var ifloat  float64 = float64(i)
}
