package main

import "fmt"

func main() {

	// a := []string{"ali", "veli"}
	// for i, k := range a {
	// 	fmt.Println(i, k)
	// }

	a := map[string]int{"ali": 76, "veli": 12, "ayse": 36}
	for k, v := range a {
		fmt.Println(k, v)
	}

}
