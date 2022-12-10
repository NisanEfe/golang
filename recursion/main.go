package main

import "fmt"

func faktor(n int) int {

	if n == 0 {
		return 1
	}
	return n * faktor(n-1)

}

func main() {

	fmt.Println(faktor(4))

}
