package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

type Cat struct {
}

func (d Dog) Speak() string {
	return "Bow Bow"
}

func (c Cat) Speak() string {
	return "meow meow"
}

func Describe(a Animal) {
	fmt.Println("Animal speaks : ", a.Speak())

	switch v := a.(type) {
	case Dog:
		fmt.Println("This is a Dog...", v.Speak())

	case Cat:
		fmt.Println("This is a Cat...", v.Speak())
	default:
		fmt.Println("This is an unknown animal", v.Speak())
	}

}

func main() {
	dog := Dog{}
	cat := Cat{}

	Describe(dog)
	Describe(cat)

}
