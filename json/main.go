package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	sampleHello()
	sampleEscapeHtml()
	sampleStruct()
	sampleMap()
}

func sampleHello() {
	fmt.Println("---------- sampleHello")
	message := []byte(`{"word": "hello world"}`)

	// stringだとエラー、何らかのinterfaceを用意する必要がある
	var out string
	err := json.Unmarshal(message, &out)
	if err != nil {
		// エラーになる
		log.Println(err)
	}

	// 空インタフェースにデコード(とはいえ、取得後が扱いにくい)
	var out2 interface{}
	err = json.Unmarshal(message, &out2)
	if err != nil {
		log.Println(err)
	}
	log.Println(out2)
}

// escape html for json
func sampleEscapeHtml() {
	fmt.Println("---------- sampleEscapeHtml")
	message := []byte(`{"message": "<div>hello world<div>"}`)
	var out bytes.Buffer
	json.HTMLEscape(&out, message)
	fmt.Println(out, out.String())
}

// json:"" の記述によりjsonへの変換時のフィールド名が変化する
//記載しない場合、フィールド名は変数名のまま。(先頭大文字もそのまま)
type Message struct {
	Id int `json:"id"`
	Word string `json:"word"`
	Suffix string `json:"suffix"`
}

// suffix を設定していないがエラーにはならない
func sampleStruct() {
	fmt.Println("---------- sampleStruct")

	// struct to json
	message := Message {Id: 1, Word: "hello world"}

	jByte, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jByte))

	// json to struct
	var message2 Message
	err = json.Unmarshal(jByte, &message2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(message2, message2.Word)

	// json to struct (miss type)
	jByte2 := []byte(`{"id": "a"}`)
	var message3 Message
	err = json.Unmarshal(jByte2, &message3)
	if err != nil {
		// 無事エラーが発生する
		log.Println(err)
	}
}

type Messages map[string]Message

func sampleMap() {
	fmt.Println("---------- sampleMap")
	// map to json
	messages := Messages {
		"moning": {Id: 1, Word: "moning"},
		"night": {Id: 2, Word: "night"},
	}

	log.Println(messages)
	jByte, err :=json.Marshal(messages)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(jByte))

	// json to map
	var out Messages
	err = json.Unmarshal(jByte, &out)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(out["night"].Word)
}

