package main

import (
	"fmt"
)

func main() {
	// Like the good old days // HL
	var myvar int // Go idiomatic design is to set default values everywhere possible
	var one, two int = 1, 2

	// With type inference // HL
	infered := 2
	anint, astring, abool := 5, "yup", true

	// Also constants // HL
	const changemeifyoucan = "Nop !"

	// Compiler won't let me keep a variable unused // HL
	fmt.Println("Myvar:", myvar, " - ", "One + Two:", one + two)
	fmt.Println("Infered:", infered, " - ", "Threes:", anint, astring, abool)
	fmt.Println("Constant:", changemeifyoucan)
}