package main

import "fmt"

func main() {

	// var a [4]string

	// a[0] = "ali"
	// a[1] = "veli"
	// a[2] = "ahmet"
	// a[3] = "mehmet"
	// fmt.Println(a)
	// fmt.Println(a[0])
	// fmt.Println(len(a))

	// a := [4]string{"ali", "ayse", "veli", "deniz"}
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }

	// a := [4]int{12, 34, 16, 43}
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i] + 3)
	// }

	a := [3]int{12, 23, 54}
	b := a
	fmt.Println(b)
	b[0] = 10
	fmt.Println(b)
	fmt.Println(a)

}
