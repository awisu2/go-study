package main

import (
	"fmt"
	"log"
	"strconv"
)

func panickIfError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	i, err := strconv.Atoi("123")
	panickIfError(err)
	fmt.Println(i) // 123

	s := strconv.Itoa(234)
	fmt.Println(s) // 234
}