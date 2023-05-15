package main

import "fmt"

type Animal interface {
	move()
}

type Food interface {
	eat() string
}

type Movement interface {
	move() string
}

type Cat struct {
	Legs int
}

func (cat *Cat) eat() string {
	return fmt.Sprintf("I am a %T and I am eating food", cat)
}

func (cat *Cat) move() string {
	return fmt.Sprintf("I am a %T and I am moving with %d legs", cat, cat.Legs)
}

func main() {

	var cat Cat = Cat{
		Legs: 4,
	}

	var food Food
	var movement Movement

	food = &cat
	movement = &cat

	fmt.Println(food.eat())
	fmt.Println(movement.move())

}
