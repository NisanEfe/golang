package main

import "fmt"

// func topla(a, b int) int {

// 	return a + b

// }

// func main() {

// 	fmt.Println(7 + 8)

// }

func islemler(a, b int) (int, int, int, int) {

	toplam := a + b
	carp := a * b
	cikar := a - b
	bol := a / b
	return toplam, carp, cikar, bol

}

func main() {

	k, l, m, n := islemler(5, 6)
	fmt.Println("toplami=", k)
	fmt.Println("carpimi=", l)
	fmt.Println("cikarma=", m)
	fmt.Println("bolumu=", n)

}
