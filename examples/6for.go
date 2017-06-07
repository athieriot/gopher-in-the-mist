package main

import "fmt"

func main() {
	sum := 0
	// Can have an init part // HL
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// On arrays // HL
	nums := []int{2, 3, 4}
	for i, v := range nums {
		fmt.Println("[", i, "]:", v)
	}

	// On maps // HL
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
