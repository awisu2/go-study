package main

import (
	"testing"
)

func TestCounter(t *testing.T) {
	callback := func(i int) {
		if i != 1 && i != 3 {
			t.Errorf("got %v. not target value", i)
		}
	}
	countUp := counter(callback)

	countUp(1)
	countUp(2)
}
