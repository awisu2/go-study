package main

import (
	"testing"
)

func TestIndexofString(t *testing.T) {
	slice := []string{"a", "b", "c"}

	if got := IndexOf(slice, "b", 0); got != 1 {
		t.Errorf("got %v. want 1", got)
	}

	if got := IndexOf(slice, "b", 2); got != -1 {
		t.Errorf("got %v. want -1", got)
	}

	if got := IndexOf(slice, "z", 0); got != -1 {
		t.Errorf("got %v. want -1", got)
	}
}

func TestIndexofInt(t *testing.T) {
	slice := []int{1, 2, 3}

	if got := IndexOf(slice, 2, 0); got != 1 {
		t.Errorf("got %v. want 1", got)
	}

	if got := IndexOf(slice, 9, 0); got != -1 {
		t.Errorf("got %v. want -1", got)
	}
}
