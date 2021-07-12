package main

import (
	"fmt"
	"log"
)

func main() {
	sampleSetValue()
	sampleSetFunction()
}

type Message struct {
	word string
}

func sampleSetValue() {
	fmt.Println("---------- sampleSetValue")

	// どんな値でもセットでき上書きも可能
	var i interface{}

	i = 1
	log.Println(i, i.(int))

	i = "a"
	log.Println(i, i.(string))

	i = Message{"hello"}
	log.Println(i, i.(Message))

	// 変換取得はエラーキャッチが可能
	i = 1
	if v, ok := i.(string); !ok  {
		log.Println("i is not string", v)
	}
}

// 共有関数の宣言
//
// ダックタイピングの分離番のような(引数の指定時にコンパイルエラー判定がされるので、分離しやすい)
// 細かい部品を増やす感じに使うと良いかも
//
// 実装内容:
// Man は getName() を持ち、それを利用するhello() が宣言されている
// MoningMan, NightMan は、それぞれ getName() を実装
// helloにManとして引き渡すことで同一の関数を実行できる
//
// NOTE: 関数しか宣言できない
// NOTE: interfaceを満たす関数が宣言されていないとエラーになる
// NOTE: クラスのような継承的な書き方はできない
//
type Man interface {
	getName() string
}

func hello(m Man) {
	fmt.Printf("good %s !\n", m.getName())
}

type MoningMan struct {
}

type NightMan struct {
}

func (m MoningMan) getName() string {
	return "moning"
}

func (m NightMan) getName() string {
	return "night"
}

type NoMan struct {
}

func sampleSetFunction() {
	fmt.Println("---------- sampleSetFunction")
	m, n := MoningMan{}, NightMan{}
	hello(m)
	hello(n)

	// getName() が実装されていないので、コンパイル時にエラーになる
	// no := NoMan{}
	// hello(no)
}
