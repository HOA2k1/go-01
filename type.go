package main

import (
	"fmt"
)

func main()  {
	// Kiểu boolean true/false
	//var a bool = true
	var a bool = false
	fmt.Println(a)

	// Kiểu string
	var str string = "hihi"
	fmt.Println(str)

	// Kiểu int
	var i int = 1
	fmt.Println(i)

	// int 8, 16, 32, 64
	// Range (trong khoảng); Range int8 -128 - 127
	// Bits


	// Kiểu unit : kiểu nguyên dương
	// Kiểu byte: alias của kiểu uint8

	// Kiểu float gồm float32 và float64
	var fl float64
	fmt.Println(fl)

	// Kiểu complex (số phức)
	var z complex64 = 10 + 1i
	fmt.Println(z)

	//Kiểu Rune: bản chất của kiểu int32
	var b string = "Nhẫn"
	runes := []rune(b)
	for i:=0; i<len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
}
