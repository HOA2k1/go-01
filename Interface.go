package main

import "fmt"

type Animal interface {
	speak()
}
type Movement interface {
	move()
}
// Embed interface
type NextAnimal interface {
	Movement
	Animal
}

//empty interface



type Dog struct {}

func (d Dog)  speak(){
	fmt.Println("gau gau")
}
func (d Dog)  move(){
	fmt.Println("đi bằng 4 chân")
}
func main()  {
	//var animal Animal
	//animal = Dog{}
	//animal.speak()
	dog := Dog{}
	var m Movement = dog
	m.move()

	var a Animal = dog
	a.speak()

	var n NextAnimal = dog
	n.move()
	n.speak()
}
