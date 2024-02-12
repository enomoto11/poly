package domain

import "github.com/google/uuid"

type cat struct {
	id   uuid.UUID
	name string
}

// NewCat returns a new cat
// Type cat implements the Animal interface
// and the signature of the Speak method
// We assign uuid to the id field so that the data store does not need to generate it
func NewCat(name string) (Animal, error) {
	cat := cat{}
	cat.id = uuid.New()
	cat.name = name

	// FIXME: no validation

	return &cat, nil
}

func (c cat) Speak() string {
	return "Meow"
}
