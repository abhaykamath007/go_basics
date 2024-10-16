// In Go, defer is a keyword that allows you to schedule a function call to be executed after the surrounding function returns.

package main

import (
	"fmt"
)

var items []string

func addItem(item string) {
	items = append(items, item)
	fmt.Println("Item added : ", item)

	defer func() {
		if len(items) > 0 {
			fmt.Println("item removed : ", item)
			items = items[:len(items)-1]
		}
	}()

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

	fmt.Println("Done inserting")

}
