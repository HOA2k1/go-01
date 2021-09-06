package main

import "fmt"

func main()  {
	// if
	a:= 1
	if a==1 {
		fmt.Println(a)
	}

	// if else
	if a>1 {
		fmt.Println(">1")
	}else {
		fmt.Println("==1")
	}


	// if statement, condition { //code}
	if b:=10; b>10{
		fmt.Println(">10")
	}else{
		fmt.Println("=10")
	}


	//switch case
	number := 5
	switch number {
	case 1:
		fmt.Println("number 1")
	case 2:
		fmt.Println("number 2")
	default:
		fmt.Println("number 10")
	}
	//fallthrough: nếu đúng thực hiện tiếp các case còn lại


	// Loops
	// for init; condition; post

	//break, continue
	for i := 0; i<10; i++ {
		if i == 4 {
			//break: chương trình sẽ thoát khỏi vòng lặp tại đây
			continue // chương trình quay lại tiếp tục chạy, bỏ qua các câu lệnh phía sau
		}
		fmt.Println(i)
	}


	// Trong go chỉ có lặp for. câu lệnh while có thể viết như sau
	j := 0
	for ;j<10;{
		fmt.Println(j)
		j++
	}

	// Chương trình lặp vô hạn (infinite loop)
	for {
		fmt.Println("infinite loop")
	}
}
