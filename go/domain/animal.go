package domain

import "errors"

type Animal interface {
	Speak() string
}

func NewAnimal(kind string, name string) (Animal, error) {
	switch kind {
	case "dog":
		return NewDog(name)
	case "cat":
		return NewCat(name)
	default:
		return nil, errors.New("unknown animal")
	}
}
