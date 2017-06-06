package main

import "fmt"

type Book struct {
	Title  string
	Review int
}

type EBook struct {
	Book
	Download string
}

type Author struct {
	Publications []Book
}

func main() {
	lestat := EBook{}
	lestat.Title = "Vampire Lestat"
	lestat.Download = "http://amazon.com"

	fmt.Println(lestat.Title)
}
