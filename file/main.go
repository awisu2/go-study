package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("tmp", 0777)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}