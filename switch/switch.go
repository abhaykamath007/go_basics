package main

import "fmt"

func main() {

	var day int
	fmt.Println("Enter the day of the week\nEnter no between 1 to 7")
	fmt.Scanln(&day)

	switch day {
	case 1:
		fmt.Println(day, " stands for Sunday")

	case 2:
		fmt.Println(day, " stands for Monday")

	case 3:
		fmt.Println(day, " stands for Tuesday")

	case 4:
		fmt.Println(day, " stands for Wednesday")

	case 5:
		fmt.Println(day, " stands for Thursday")

	case 6:
		fmt.Println(day, " stands for Friday")

	case 7:
		fmt.Println(day, " stands for Saturday")

	default:
		fmt.Println(day, "invalid ")
	}
}
