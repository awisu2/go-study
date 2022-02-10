package main

import (
	"testing"
)


func TestAddCodeEachWord(t *testing.T) {
    got := AddCodeEachWord("aあ1", 1)
	if got != "bぃ2" {
		t.Errorf("got %v. want bぃ2", got)
	}
}
