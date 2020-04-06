package main

import "go_datatypes/organization"

func main() {
	p := organization.NewPerson("Darth","Vader")
	println(p.ID())
}
