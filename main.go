package main

import (
	"fmt"
	"go_datatypes/organization"
)

func main() {
	p := organization.NewPerson("Darth", "Vader")
	println(p.ID())
	println(p.FullName())
	err := p.SetTwitterHandler("darth")
	if err != nil {
		fmt.Printf("An error occured setting twitter hander: %s", err.Error())
	}

}
