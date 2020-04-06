package main

import "go_datatypes/organization"

func main() {
	p := organization.Person{FirstName: "Darth", LastName: "Vader"}
	println(p.ID())
}
