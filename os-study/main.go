package main

import (
	"io/ioutil"
	"os"
)

// read any function's stdout
func ReadStdOut(f func()) ([]byte, error) {
	stdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout

	return ioutil.ReadAll(r)
}
