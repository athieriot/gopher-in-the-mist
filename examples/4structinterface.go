package main

import "fmt"

type EBook struct {
	Title  string
	Review int
}

type Booker interface {
	IsItGood() bool 	// Public
	hasInStock() bool 	// Private
}

func (b EBook) IsItGood() bool { return b.Review > 5 }
func (b EBook) hasInStock() bool { return false }

func main() {
	lestat := EBook{}
	lestat.Title = "Vampire Lestat"
	lestat.Review = 10

	fmt.Println(lestat.IsItGood())
	fmt.Println(Booker(lestat).hasInStock())
}
