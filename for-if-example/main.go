package main

import "fmt"

func main() {

	a := 1000
	toplam := 0
	for i := 0; i < a; i++ {
		if i%3 == 0 || i%5 == 0 {
			toplam += i
		}
	}
	fmt.Println(toplam)

}
