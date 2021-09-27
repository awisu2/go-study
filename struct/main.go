package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	A string `json:"a", foo:"bar"`
	B string `json:"b"`
}

func main() {
	tagInfo()
}

func tagInfo () {
	sample := Sample{A: "aaa", B: "bbb"}

	// reflectによりフィールド情報を取得
	field, ok := reflect.TypeOf(&sample).Elem().FieldByName("A")
	if !ok {
		panic("field not found")
	}

	fmt.Println(field.Tag) // json:"a", foo:"bar"
}