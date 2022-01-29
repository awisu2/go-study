package main

import (
	"testing"
)

func TestReplace(t *testing.T) {
	got := replace("abcdddddddefg", `[d]*ef`, "")
	if got != "abcg" {
		t.Errorf("got %v. want abcg", got)
	}
}