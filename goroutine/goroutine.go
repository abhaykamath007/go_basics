package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printNumbers(prefix string, no int) {

	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Printf("%s%d: %d\n", prefix, no, i)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("done ", prefix, no)
}

func main() {
	start := time.Now()
	fmt.Println("Starting goroutines...")

	defer func() {
		fmt.Println(time.Since(start))
	}()

	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go printNumbers("Goroutine", i)
	}

	fmt.Println("Main function completed.")
	wg.Wait()
}
