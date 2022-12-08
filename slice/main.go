package main

import "fmt"

func main() {

	// var a []int
	// fmt.Println(a)

	// a := make([]int, 3)
	// fmt.Println(a)

	// var a []int
	// a = append(a, 12, 35, 68)
	// fmt.Println(a)

	// a := []int{12, 34}
	// a = append(a, 45, 67)
	// fmt.Println(a)

	// a := []int{12, 23, 54}
	// b := a
	// fmt.Println(b)
	// b[0] = 10
	// fmt.Println(b)
	// fmt.Println(a)
	// slice -> referanslari kopyaliyoruz

	// a := []int{12, 23, 54, 56, 78}
	// fmt.Println(a[4])
	// fmt.Println(a[2:4])
	// fmt.Println(a[:3])
	// fmt.Println(a[2:])

	// a := []int{12, 23, 54, 56, 78}
	// b := a[2:4]
	// fmt.Println(b)

	// a := []int{12, 23, 54, 56, 78}
	// b := a[2:4]
	// b[0] = 200
	// fmt.Println(b)
	// fmt.Println(a)

	a := []int{12, 23, 54, 56, 78}
	b := make([]int, len(a))
	copy(b, a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("---")
	b[0] = 199
	fmt.Println(a)
	// kopyalanan slicelar da degerler etkilenmez

}
