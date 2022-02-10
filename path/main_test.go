package main

import (
	"testing"
)


func TestRenameWindows(t *testing.T) {
    got := ReBaseName("c:/a/b/c.jpg", "d", ReBaseNameOption{})
	if got != `c:\a\b\d.jpg` {
		t.Errorf(`got %v. want c:\a\b\d.jpg`, got)
	}

	got = ReBaseName("c:/a/b/c.jpg", "d", ReBaseNameOption{
		IsDirectory: true,
	})
	if got != `c:\a\b\d` {
		t.Errorf(`%v. want c:/a/b/d`, got)
	}
}

func TestSplitName(t *testing.T) {
	name, ext := SplitName("abc.d.jpg")
	if name != "abc.d" {
		t.Errorf("%v. want abc.d", name)
	}
	if ext != ".jpg" {
		t.Errorf("%v. want .jpg", ext)
	}
}

func TestSplitNameNoExt(t *testing.T) {
	name, ext := SplitName("abc")
	if name != "abc" {
		t.Errorf("%v. want abc", name)
	}
	if ext != "" {
		t.Errorf("%v. want ''", ext)
	}
}