package main

import "fmt"

func main()  {
	a:= 100.1
	var pointer *float64 // zero value nil
	pointer = &a
	fmt.Println(pointer)

	// cách 2
	b := 123
	p2 := new(int)  // zero value không nil
	p2 = &b
	fmt.Println(p2)

	// *p2 = 99 => b sẽ bị thay đổi giá trị thành 99


	// con trỏ trỏ tới array
	aray := [3] int {1,2,3}
	var p *[3]int
	p = &aray
	fmt.Println(p)

}