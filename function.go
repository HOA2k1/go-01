package main

import "fmt"

// Khai báo function: func func_name (params) return_type { //code}

func Hello() {
	fmt.Println("Xin chào")
}

func sayHello(name string)  {
	fmt.Println("Xin chào:", name)
}

func greeting(name string) string  {
	resulf := fmt.Sprintf("Hello %s", name)
	return resulf
}

// Multiple return values
func number(a, b int)  (int, int){
	return a, b
}

// Named return values
func equal1(a1, b1 int)  (a int, b int, equal bool){
	equal = a1 == b1
	return a1, b1, equal
}

//Variadic Functions
func addItem(item int, list ...int)  {
	list = append(list, item)
	fmt.Println(list)
}

func change(list ...int)  {
	list[0] = 21
}
func main()  {
	Hello()
	sayHello("Hoa")
	resulf := greeting("hoa")
	fmt.Println(resulf)
	a, b := number(10, 20)
	fmt.Println("a = ",a," b = ",b)
	a1,b1,equal := equal1(20,10)
	if equal{
		fmt.Println("Bang")
	}else {
		fmt.Println("a= ",a1," b = ",b1)
	}


	//Variadic Functions
	addItem(1, 100,200,300,400)
	var list = []int {1,2,3,4}
	addItem(100, list...)

	change(list...)
	fmt.Println(list)
}