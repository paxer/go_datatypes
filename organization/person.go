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

type Person struct {
	firstName      string
	lastName       string
	twitterHandler TwitterHandler
}

func NewPerson(firstName, lastName string) Person {
	return Person{firstName: firstName, lastName: lastName}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p *Person) ID() string {
	return "12345"
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
