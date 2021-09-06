package main

import "fmt"

type Student struct {
	id int
	name string
}
func main()  {
	// cách 1
	st1 := Student{
		id: 123,
		name: "Bong",
	}
	fmt.Println(st1)

	//Cách 2
	st2 := Student{123,"Hoa"}
	fmt.Println(st2.name)

	//cách 3
	var st3 Student= struct {
		id   int
		name string
	}{id: 2222, name: "hihi"}
	fmt.Println(st3)


	//anonymous struct: struct vô danh
	var anonymous = struct {
		email string
		age int
	}{
		"hoa@gmail.com",
		20,
	}
	fmt.Println(anonymous)


	// pointer trỏ tới struct
	st4 := &Student{
		21,
		"Do",
	}
	fmt.Println(&st4)
	fmt.Println(st4.id)
}