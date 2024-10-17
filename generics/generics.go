//

package main

import (
	"fmt"
)

type List[T comparable] struct {
	array []T
}

func (l *List[T]) Add(val T) {
	l.array = append(l.array, val)
}

func (l *List[T]) Remove(val T) bool {

	for i, value := range l.array {
		if value == val {
			l.array = append(l.array[:i], l.array[i+1:]...)
			return true
		}
	}
	return false
}

func (l List[T]) Display() {
	fmt.Println("Values : ")
	for _, value := range l.array {
		fmt.Print(value, " ")
	}
}

func multiply[T comparable](a T, b T) {
	fmt.Println(a, b)
}

func main() {

	intList := List[int]{}

	intList.Add(23)
	intList.Add(34)
	intList.Add(78)
	intList.Add(790)
	intList.Add(234)

	fmt.Println("Displaying the elements : ")
	fmt.Scanln()
	intList.Display()

	fmt.Println("\nRemoving elements :")
	res := intList.Remove(34)
	if res {
		fmt.Println("Successfully removed the element")
	} else {
		fmt.Println("failed")
	}
	fmt.Scanln()

	res = intList.Remove(482787)
	if res {
		fmt.Println("Successfully removed the element")
	} else {
		fmt.Println("failed")
	}

	fmt.Println("After removing .....")

	intList.Display()

	fmt.Scanln()

	fmt.Println("......Strings.......")

	stringList := List[string]{}

	stringList.Add("Abhay")
	stringList.Add("Ajit")
	stringList.Add("Akshay")

	stringList.Display()

	stringList.Remove("Ajit")
	stringList.Display()

	multiply(13, 14)

}
