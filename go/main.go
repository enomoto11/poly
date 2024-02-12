package main

import (
	"fmt"

	"github.com/enomoto11/poly/domain"
)

func main() {
	// Create a new dog
	dog, err := domain.NewAnimal("dog", "koro")
	if err != nil {
		fmt.Println(err)
	}

	// Create a loyal dog
	loyalDog, err := domain.NewAnimal("dog", "pochi")
	if err != nil {
		fmt.Println(err)
	}

	// Create a new cat
	cat, err := domain.NewAnimal("cat", "tama")
	if err != nil {
		fmt.Println(err)
		return
	}

	res1 := dog.Speak()
	res2 := loyalDog.Speak()
	res3 := cat.Speak()

	fmt.Println(res1) // Woof!
	fmt.Println(res2) // I'm loyal! I'm pochi!
	fmt.Println(res3) // Meow
}
