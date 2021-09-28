package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	A string `json:"a", foo:"bar"`
	B string `json:"b"`
}

// Sampleを直接宣言し、同じフィールドを設定した場合
type Sample2 struct {
	Sample
	B string `json:"b"`
}

func main() {
	tagInfo()
	override()
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

func override() {
	s2 := Sample2 {
		// セット時、直接のアクセスはできない
		// A: "a",
		Sample: Sample{
			A: "a",
			B: "b1",
		},
		B: "b2",
	}

	fmt.Println(s2) // {{a b1} b2}

	// 宣言後は直接のアクセスが可能
	s2.A = "aaa"
	s2.Sample.B = "bbb1"
	s2.B = "bbb2"

	// 取得時も同じく直接アクセス可能
	fmt.Println(s2.A) // aaa
	// 元になったstructのプロパティが優先される
	fmt.Println(s2.B) // bbb2
}