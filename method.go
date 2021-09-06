package main

import "fmt"

type Sinhvien struct {
	name string
}

//Define method
// func (t Type) methodName(params) return {//code here}
// (t Type) => receiver
func (s Sinhvien) getName()  string{
	return s.name
}

// Value receiver: không bị thay đổi giá trị trong struct (không ảnh hưởng tới những cái ban đầu)
func (s Sinhvien) changeName() {
	s.name = "An"
}

// Pointer receiver:
func (s *Sinhvien) changeName2() {
	s.name = "An"
}
func main()  {
	st := Sinhvien{"hoa"}
	st.changeName2()
	fmt.Println(st.name)
}
