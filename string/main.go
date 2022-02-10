package main

import (
	"fmt"
	"strings"
)

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
	s := "xaybx"
	for i, code := range "abc" {
		fmt.Printf("%d: %c\n", i, code)
		s = strings.ReplaceAll(s, string(code), "=")
	}
	fmt.Println(s) // x=y=z


	fmt.Println(strings.Replace("aaaa", "a", "b", 2)) // bbaa
	fmt.Println(strings.Replace("aaaa", "a", "b", -1)) // bbbb
}

func main() {
	sample()
}