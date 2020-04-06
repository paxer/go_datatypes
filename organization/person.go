package organization

import (
	"errors"
	"fmt"
	"strings"
)

type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type socialSecurityNumber string

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "Unites States of America"
}

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

type europeanUnionIdentifier struct {
	id string
	country string
}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeanUnionIdentifier) Country() string {
	return eui.country
}

func NewEuropeanUnionIdentifier(id, country string) Citizen {
	return europeanUnionIdentifier{id: id, country: country}
}


type Name struct {
	first string
	last string
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}

type Person struct {
	Name
	twitterHandler TwitterHandler
	Citizen
}

func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			first: firstName,
			last: lastName,
		},
		Citizen: citizen,
	}
}


func (p *Person) ID() string {
	return fmt.Sprintf("Person's idendtifier: %s", p.Citizen.ID())
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
