package main

import (
	"os"
	"testing"
)

func TestMkdir(t *testing.T) {
	const dir = "tmp"
	err := os.RemoveAll(dir)
	if err != nil {
		t.Errorf("%v", err)
	}

	err = os.Mkdir(dir, 0777)
	if err != nil {
		t.Errorf("%v", err)
	}

	err = os.Mkdir(dir, 0777)
	if err == nil {
		t.Errorf("no error. want already exists error")
	}

	err = os.Remove(dir)
	if err != nil {
		t.Errorf("%v", err)
	}

	err = os.Remove(dir)
	if err == nil {
		t.Errorf("no error. want not exists error")
	}
}
