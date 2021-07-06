package main

import (
	"fmt"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
    fmt.Println("Hello, World!")

		fmt.Println(quote.Go())

		message := greetings.Hello("foo")
		fmt.Println(message)
}
