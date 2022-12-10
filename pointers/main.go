package main

import "fmt"

func sent(t *int) {

	fmt.Println(t)
	*t = 32

}

func main() {

	a := 67
	fmt.Println(&a)
	sent(&a)
	fmt.Println(a)

}

// func main() {

// 	a := 23
// 	fmt.Println(a)
// 	fmt.Println(&a)
// 	b := &a
// 	fmt.Println(b)
// 	fmt.Println(*b)
// 	*b = 45
// 	fmt.Println(a)

// }
