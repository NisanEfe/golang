package main

import "fmt"

type contactInterface interface {
	save()
	update()
	delete()
}

type contact struct {
	personName   string
	personNumber string
}

func imple(k contactInterface) {
	k.save()
}

func (r contact) delete() {
	fmt.Println(r.personName, "delete")
}

func (r contact) save() {
	fmt.Println(r.personName, "save")
}

func (r contact) update() {
	fmt.Println(r.personName, "update")
}

func main() {

	personObject := contact{personName: "ali", personNumber: "123"}
	fmt.Println(personObject)
	imple(personObject)

}
