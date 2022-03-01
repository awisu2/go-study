package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestTypeOfString(t *testing.T) {
	s := "hello"
	i := interface{}(s)

	// get Type: [Type](https://pkg.go.dev/reflect#Type)
	typ := reflect.TypeOf(i)

	// [Kind](https://pkg.go.dev/reflect#Kind)
	if typ.Kind() != reflect.String {
		t.Errorf("Kind() is not reflect.String")
	}

	if got := typ.String(); got != "string" {
		t.Errorf("got %v. want string", got)
	}
	if got := typ.Name(); got != "string" {
		t.Errorf("got %v. want string", got)
	}
}

func TestTypeOfSlice(t *testing.T) {
	v := []int{1, 2, 3}
	i := interface{}(v)

	// get Type: [Type](https://pkg.go.dev/reflect#Type)
	typ := reflect.TypeOf(i)

	// [Kind](https://pkg.go.dev/reflect#Kind)
	if typ.Kind() != reflect.Slice {
		t.Errorf("Kind() is not reflect.Slice")
	}

	if got := typ.String(); got != "[]int" {
		t.Errorf("got %v. want []int", got)
	}
	if got := typ.Name(); got != "" {
		t.Errorf("got '%v' want ''", got)
	}

	// get element type
	elmType := typ.Elem()
	if got := elmType.Kind(); got != reflect.Int {
		t.Errorf("got %v. want reflect.Int", got)
	}
}

func TestWithStruct(t *testing.T) {
	type User struct {
		Name   string
		Length float32
	}
	v := User{
		Name:   "jon",
		Length: 156.78,
	}
	i := interface{}(v)
	rf := reflect.TypeOf(i)

	names := []string{}
	types := []reflect.Type{}
	for i := 0; i < rf.NumField(); i++ {
		f := rf.Field(i)
		names = append(names, f.Name)
		types = append(types, f.Type)
	}
	if names[0] != "Name" || names[1] != "Length" {
		t.Errorf(`got %v. want ["Name", "Length"]`, names)
	}
	if types[0].Kind() != reflect.String || types[1].Kind() != reflect.Float32 {
		t.Errorf(`got %v. want [string, float32]`, names)
	}
}

func TestPrintAnyValue(t *testing.T) {
	if got, _ := readStdOut(func() {
		PrintAnyValue(reflect.ValueOf(1))
	}); string(got) != "1\n" {
		t.Errorf(`got '%v'. want 1`, string(got))
	}

	if got, _ := readStdOut(func() {
		PrintAnyValue(reflect.ValueOf("a"))
	}); string(got) != "a\n" {
		t.Errorf(`got '%v'. want a`, string(got))
	}
}

func readStdOut(f func()) ([]byte, error) {
	stdout := os.Stdout
	defer func() {
		os.Stdout = stdout
	}()

	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()

	return ioutil.ReadAll(r)
}

func TestValue(t *testing.T) {
	v := reflect.ValueOf("a")
	if v.Kind() != reflect.String {
		t.Error("not string")
	}
	if v.String() != "a" {
		t.Error("not a")
	}
}
