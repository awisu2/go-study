package main

import (
	"fmt"
	"reflect"
)

// check Elem() is available
func HasElem(rf reflect.Type) bool {
	switch rf.Kind() {
	case reflect.Array:
		return true
	case reflect.Chan:
		return true
	case reflect.Map:
		return true
	case reflect.Ptr:
		return true
	case reflect.Slice:
		return true
	default:
		return false
	}
}

// v can get any type
func PrintAnyValue(v reflect.Value) {
	fmt.Println(v)
}
