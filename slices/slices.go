package main

import "fmt"

func main() {
	var fruit string
	fruits := []string{}

	fmt.Println("Enter the names of the fruits : ")

	for {
		fmt.Println("Enter the fruit name : ")
		fmt.Scanln(&fruit)

		if fruit == "done" || fruit == "" {
			break
		}

		fruits = append(fruits, fruit)
	}

	if len(fruits) == 0 {
		fmt.Println("There are no fruits available")

	} else {

		fmt.Println("The fruits are : ")
		for _, fruit := range fruits {
			fmt.Println(fruit)
		}
	}
}
