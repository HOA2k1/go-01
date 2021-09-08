package main

import ("fmt"
	"strings"
)
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	IPAddrString := []string{}
	//lặp các phần tử lấy value
	for _, value := range ip {
		IPAddrString = append(IPAddrString, fmt.Sprint(int(value))) //mỗi lần lặp thì thêm 1 phần tử vào IPAddrString
	}
	return strings.Join(IPAddrString, ".") // Nối các phần tử thành 1 chuỗi, phân tách sep đặt giữa các phần tử
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8},
	}



	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
