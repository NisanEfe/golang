package main

import "fmt"

func main() {

	// c := map[string]int{"ali": 76, "veli": 12}

	// a := make(map[string]int)
	// a["ali"] = 76
	// a["veli"] = 12
	// fmt.Println(a)

	// var a map[int]string
	// fmt.Println(a)

	// a := map[string]int{"ali": 76, "veli": 12, "ayse": 26}
	// fmt.Println(a["veli"])
	// fmt.Println(len(a))
	// delete(a, "veli")
	// fmt.Println(a)

	// a := map[string]int{"ali": 76, "veli": 12, "ayse": 26}
	// fmt.Println(a)
	// delete(a, "veli")
	// fmt.Println(a)

	// a := map[string]int{"ali": 76, "veli": 12, "ayse": 26}
	// k, m := a["ali"]
	// fmt.Println(k, m)

	a := map[string]int{"ali": 76, "veli": 12, "ayse": 26}
	_, m := a["kemal"]
	fmt.Println(m)

}
