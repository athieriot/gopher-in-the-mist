package main

import (
	"fmt"
)

func player(messages chan string, name string) {
	fmt.Println(name, "ping")
	messages <- name
}

func main() {
	messages := make(chan string)

	go player(messages, "Player 1")
	go player(messages, "Player 2")

	msg := <-messages
	fmt.Println(msg, "pong")

	msg = <-messages
	fmt.Println(msg, "pong")

	close(messages)
}
