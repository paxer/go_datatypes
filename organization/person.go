package organization

import "fmt"

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName string
	lastName  string
}

func NewPerson(firstName, lastName string) Person {
	return Person{firstName: firstName, lastName: lastName}
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p Person) ID() string {
	return "12345"
}
