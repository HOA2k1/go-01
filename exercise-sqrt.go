package main

import (
	"fmt"
)

const Delta = 0.0001

func isConverged(d float64) bool {
	if d < 0.0 {
		d = -d
	}
	if d < 0.0001 {
		return true
	}
	return false
}

func Sqrt(x float64) float64 {
	z := 1.0
	tmp := 0.0
	for {
		tmp = z - (z * z - x) / 2 * z
		if d := z - tmp; isConverged(d) {
			return tmp
		}
		z = tmp
	}
	return z
}

func main() {
	a := Sqrt(2)
	fmt.Println(a)
}