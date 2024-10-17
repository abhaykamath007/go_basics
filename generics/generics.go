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

func main() {

	intList := List[int]{}
	stringList := List[string]{}

	var choice int

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add Integer")
		fmt.Println("2. Remove Integer")
		fmt.Println("3. Display Integer List")
		fmt.Println("4. Add String")
		fmt.Println("5. Remove String")
		fmt.Println("6. Display String List")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var val int
			fmt.Print("Enter an integer to add: ")
			fmt.Scan(&val)
			intList.Add(val)
			fmt.Println("Added:", val)

		case 2:
			var val int
			fmt.Print("Enter an integer to remove: ")
			fmt.Scan(&val)
			if intList.Remove(val) {
				fmt.Println("Successfully removed:", val)
			} else {
				fmt.Println("Failed to remove:", val)
			}

		case 3:
			fmt.Println("Displaying Integer List:")
			intList.Display()

		case 4:
			var val string
			fmt.Print("Enter a string to add: ")
			fmt.Scan(&val)
			stringList.Add(val)
			fmt.Println("Added:", val)

		case 5:
			var val string
			fmt.Print("Enter a string to remove: ")
			fmt.Scan(&val)
			if stringList.Remove(val) {
				fmt.Println("Successfully removed:", val)
			} else {
				fmt.Println("Failed to remove:", val)
			}

		case 6:
			fmt.Println("Displaying String List:")
			stringList.Display()

		case 7:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}
