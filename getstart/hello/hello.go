package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
		log.SetPrefix("greetings: ")
		log.SetFlags(0)

		fmt.Println("Hello, World!")

		fmt.Println(quote.Go())

		// get hello
		message, err := greetings.Hello("foo")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(message)

		// get hellos
		names := []string{"foo", "bar", "smith", "neo"}
		messages, err := greetings.Hellos(names)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(messages)

		// get error
		message2, err := greetings.Hello("")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(message2)
}
