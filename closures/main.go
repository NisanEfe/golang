package main

import "fmt"

// func person(name string) func() string {

// 	_name := name
// 	persons := func() string {
// 		fmt.Println(_name)
// 		return "geri donen " + _name
// 	}
// 	return persons
// }

// func main() {

// 	newPerson := person("efe")
// 	newPerson()
// 	a := newPerson()
// 	fmt.Println(a)

// }

func rating() (birinci func(), lastCount func() int) {

	count := 0
	birinci = func() {
		count++
		fmt.Println(count)
	}

	lastCount = func() int {
		return count
	}
	return
}

func main() {

	ratingCount, lastCount := rating()
	for i := 0; i < 10; i++ {
		ratingCount()
	}
	fmt.Println(lastCount())

}
