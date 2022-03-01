package main

import (
	"io/ioutil"
	"os"
)

// read any function's stdout
func ReadStdOut(f func()) ([]byte, error) {
	stdout := os.Stdout
	defer func() {
		os.Stdout = stdout
	}()

	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()

	return ioutil.ReadAll(r)
}
