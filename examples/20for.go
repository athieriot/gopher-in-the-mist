package main

import "fmt"

func main() {
	sum := 0
	// With an init part // HL
	for i := 0; i < 10; i++ {
		sum += i
	}
	sum = 0
	// Init optional // HL
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
