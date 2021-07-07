package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }

	// `:=` は `var message string` と値のセットを同時に行っている
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {
    messages := make(map[string]string)

    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        messages[name] = message
    }

    return messages, nil
}

// initialize時にrunされる関数(多分import時に実行され、main関数より早く実行される)
// [Effective Go - The Go Programming Language](https://golang.org/doc/effective_go#init)
func init () {
    log.Println("greeting's init")
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
     return formats[rand.Intn(len(formats))]
}