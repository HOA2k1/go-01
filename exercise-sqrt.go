package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	t := 0.0	
    for {
        z, t = z - (z*z-x)/(2*z), z
        if math.Abs(t-z) < 0.00000001 {
            break
        }
    }
    return z
}

func main() {
	a := Sqrt(3.1)
	fmt.Println(a)	
}