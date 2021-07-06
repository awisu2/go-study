package greetings

import "fmt"

func Hello(name string) string {
	// `:=` は `var message string` と値のセットを同時に行っている
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
