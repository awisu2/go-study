package main

import "fmt"

// Human has human's variable
type Human struct {
	Name string // name
	// age
	// from 0 to 200
	Age string
}

// Hello print hello with name
func Hello(name string) string{
	return fmt.Sprintf("hello %s\n", name)
}

// Hi print hi with name
//
// name string
func Hi(name string) string{
	return fmt.Sprintf("hi %s\n", name)
}
