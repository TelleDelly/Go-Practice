package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Martin")
	fmt.Println(message)
}