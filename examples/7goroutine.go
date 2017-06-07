package main

import (
	"fmt"
	"time"
)

func printStuff(as string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(as, ":", i)
	}
}

func main() {
	// Blocking call // HL
	printStuff("direct")

	// Async call // HL
	go printStuff("goroutine")
	go func() { printStuff("Anonymous") }()

	time.Sleep(1 * time.Minute)
}
