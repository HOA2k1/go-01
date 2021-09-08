package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var img [][]uint8
	for i := 0; i < dy; i ++ {
		row := make([]uint8, dx)
		for j := 0; j < dx; j ++ {
			row[j] = uint8(i*j)
		}
		img = append(img, row)
	}
	return img
}

func main() {
	pic.Show(Pic)
}
