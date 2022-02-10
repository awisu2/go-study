package main

import "fmt"

// それぞれの文字コードに加算をして返却
func AddCodeEachWord(input string, add rune) string {
	out := ""
	for _, code := range input {
		out += string(code + add)
	}
	return out
}

// any sample
func sample() {
	// substring and getting code
	fmt.Println("abc") // abc

	fmt.Println("abc"[1]) // 98
	fmt.Println(string("abc"[1])) // b
	fmt.Println("abc"[1:3]) // bc

	// code print
	for i, code := range "abc" {
		fmt.Printf("%d: %v\n", i, code)
	}
}

func main() {
	sample()
}