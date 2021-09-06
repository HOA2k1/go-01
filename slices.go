package main

import "fmt"

func main()  {
	// Khai báo và khởi tạo giống như array

	// Tạo slices từ array
	var array = [5] int {1,2,3,4,5}
	//slices1 := array[:]     Lấy toàn bộ giá trị array
	//slices1 := array[:3]     Lấy giá trị từ 0 đến 2
	//slices1 := array[2:]     Lấy giá trị từ 2 đến cuối
	slices1 := array[1:3]
	fmt.Println(slices1)


	// len: số lượng phần tử trong slices
	// cap: số lượng phần tử tình từ phần tử slices đầu tiên đến hết cái array ban đầu
	var contries = [...] string {"vn", "england", "usa", "japan"}
	slices2 := contries[2:3]
	fmt.Println(len(slices2)) // lấy ra được giá trị usa trong contries => độ dài là 1
	fmt.Println(cap(slices2)) // tính từ vị trí của usa trong contries đến hết, gồm usa và japan => độ dài là 2

	//make
	slices3 := make([]int, 3, 5)
	fmt.Println(slices3)

	//append: thêm 1 giá trị vào slices
	slices4 := make([]int, 2)
	slices4 = append(slices4, 10)
	fmt.Println(slices4)

	// copy
	src := []string {"a", "b", "c", "d"}
	desc := make([]string, 2)
	copy(desc, src)
	fmt.Println(desc)

	//delete
	src = append(src[:1], src[2:]...)  // nối 2 slices với nhau = append(slice1, slice2...)

	//SliceTricks
}