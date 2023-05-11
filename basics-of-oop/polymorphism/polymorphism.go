package main

import "fmt"

type Animal interface {
	MakeSound()
}

type Dog struct{}
type Cat struct{}

func (dog *Dog) MakeSound() {
	fmt.Println("Woof")
}

func (cat *Cat) MakeSound() {
	fmt.Println("Meow")
}

func main() {
	var animals []Animal
	var animal Animal

	var dog Dog = Dog{}
	var cat Cat = Cat{}

	animal = &dog
	animals = append(animals, animal)

	animal = &cat
	animals = append(animals, animal)

	// we can get classess with accordance to the sound it makes
	for _, animal := range animals {
		animal.MakeSound()
	}
}
