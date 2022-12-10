package main

import (
	"errors"
	"fmt"
)

func checkUserName(v string) error {

	if v == "nisanefe" {
		return errors.New("bu kullanici adi alinmis")
	} else {
		return nil
	}

}

func main() {

	userNames := []string{"gurkanaydn", "nisanefe", "alierbey"}
	for _, i := range userNames {
		if a := checkUserName(i); a != nil {
			fmt.Println(a)
		} else {
			fmt.Println(i, "Kullanici adini alabilirsiniz")
		}
	}

}
