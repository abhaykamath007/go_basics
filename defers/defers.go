// In Go, defer is a keyword that allows you to schedule a function call to be executed after the surrounding function returns.

package main

import (
	"fmt"
)

var items []string

func addItem(item string) {
	items = append(items, item)
	fmt.Println("Item added : ", item)

}

func main() {
	fmt.Println("Starting...")

	defer func() {
		fmt.Println("Closing the program")
		fmt.Println("Final Items", items)
	}()

	addItem("item1")
	addItem("item2")
	addItem("item3")

	defer func() {
		for len(items) > 0 {
			if len(items) > 0 {
				items = items[:len(items)-1]
			}
			fmt.Println(items)

			for i := 1; i <= 3; i++ {
				defer fmt.Println(i)
			}
		}
	}()

	fmt.Println("Done inserting")

}
