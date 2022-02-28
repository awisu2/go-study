package main

import (
	"runtime"
	"testing"
)

// check is contain
func contain(slice []string, search string) bool {
	for _, v := range slice {
		if v == search {
			return true
		}
	}
	return false
}

func TestGOOS(t *testing.T) {
	goos := runtime.GOOS
	if !contain(GOOS_LIST, goos) {
		t.Errorf("goos is out of values. got %s", goos)
	}
}

func TestGOARCH(t *testing.T) {
	goarch := runtime.GOARCH
	if !contain(GOARCH_LIST, goarch) {
		t.Errorf("goos is out of values. got %s", goarch)
	}
}
