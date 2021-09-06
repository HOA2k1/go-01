package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci(s int) int {
	if s==0 || s ==1{
		return 1
	}else {
		return fibonacci(s-2)+fibonacci(s-1)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}
}
