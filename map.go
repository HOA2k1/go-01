package main

import "fmt"

func main()  {
	// Khai báo có 2 cách
	var myMap = make(map[string]int) // Khi này myMap != nil vì sử dụng make thì nó có 1 vùng nhớ rồi
	var myMap1 map[string]int // Khi này thì myMap1 = nil

	fmt.Println(myMap)
	fmt.Println(myMap1)


	// Khai báo có giá trị khởi tạo
	myMap3 := map[string]int{
		"key1" : 1,
		"key2" : 2,
		"key3" : 4,
	}

	//Thêm một phần tử vào map
	myMap3["key4"] = 5
	myMap3["key5"] = 5


	// delete một phần tử trong map
	delete(myMap3, "key4")
	fmt.Println(myMap3)

	//length của map: len()

	// map là reference type
	myMap2 := myMap3
	fmt.Println(myMap2)

	myMap3["key5"] = 1000  // khi thay đổi mymap3 thì mymap2 cũng thay đổi theo
	fmt.Println(myMap2)


	// truy cập 1 phần tử trong map
	value := myMap3["key2"]
	fmt.Println(value)

	// trong map không có toán tử ==
}
