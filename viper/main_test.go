package main

import (
	"testing"
)

func TestSampleSetDefault(t *testing.T) {
	set := "myvalue"
	got := sampleSetDefault(set)
	if got != set {
		t.Errorf("got %v; want %v", got, set)
	}
}

func TestSampleReadFile(t *testing.T) {
	got := sampleReadFile()
	checkValue := got.ContentDir
	if checkValue != "sampletext" {
		t.Errorf("got => '%v', want 'sampletext'", checkValue)
	}
}

func TestSampleREadEnv(t *testing.T) {
	t.Setenv("B", "abcd")
	got := sampleReadEnv()
	if got != "abcd" {
		t.Errorf("got => %v, want 'abcd'", got)
	}
}

func TestSampleREadEnvAndConfig(t *testing.T) {
	t.Setenv("B", "abcd")
	got := sampleReadEnvAndConfig()
	if got != "abcd" {
		t.Errorf("got => %v, want 'abcd'", got)
	}
}
