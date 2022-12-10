package main

import "fmt"

type book struct {
	bookName   string
	bookWriter string
	bookPage   int
}

func main() {

	// fmt.Println(book{"Denemeler", "Montaigne", 185})
	// fmt.Println(book{bookName: "Ruhi Mucerret", bookWriter: "Murat Mentes", bookPage: 280})
	// fmt.Println(book{bookName: "Hayatin Kaynagi", bookWriter: "Ayn Rand"})
	a := book{bookName: "Savas ve Baris", bookWriter: "Tolstoy", bookPage: 340}
	fmt.Println(a)
	fmt.Println(a.bookName)

}
