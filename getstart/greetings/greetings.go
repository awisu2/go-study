package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }

	// `:=` は `var message string` と値のセットを同時に行っている
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
