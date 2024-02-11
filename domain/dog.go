package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type dog struct {
	id   uuid.UUID
	name string
}

// NewDog returns a new dog
// Type dog implements the Animal interface
// We assign uuid to the id field so that the data store does not need to generate it
func NewDog(name string) (Animal, error) {
	dog := dog{}
	dog.id = uuid.New()
	dog.name = name

	if err := dog.validate(); err != nil {
		return nil, err
	}

	return &dog, nil
}

// capsulate the validation logic
func (d *dog) validate() error {
	if d.name == "" {
		return fmt.Errorf("name is required")
	}

	if len(d.name) > 50 {
		return fmt.Errorf("name must be less than 50 characters")
	}

	return nil
}

func (d *dog) Speak() string {
	if d.name == "pochi" {
		return fmt.Sprintln("I'm loyal! I'm", d.name, "!")
	}

	return "Woof!"
}
