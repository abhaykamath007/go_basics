package main

import (
	"fmt"
)

func main() {

	var age int
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)

	if age < 0 {
		fmt.Println("Invalid age  ")
	} else {
		if age < 13 {
			fmt.Println("You are a child")
		} else if age < 18 {
			fmt.Println("You are a teenager")
		} else if age < 60 {
			fmt.Println("You are an adult")
		} else {
			fmt.Println("You are a senior citizen")
		}
	}

}
