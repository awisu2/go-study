package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	sampleHello()
	sampleEscapeHtml()
	sampleStruct()
	sampleMap()
	sampleCustom()
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

	// mapでざっくり受け取る、上記の空インタフェースよりはkeyがわかるので少しマシ？
	var out3 map[string]interface{}
	d := json.NewDecoder(bytes.NewReader(message))
	// Unmarshalとの違いとして、decode前にカスタムできる
	d.DisallowUnknownFields()
	if err := d.Decode(&out3); err != nil {
		log.Println(err)
	}
	log.Println(out3)
	for k, v := range out3 {
		log.Printf("%s: %v\n", k, v)
	}
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

// json変換をカスタマイズする
//
// https://dena.github.io/codelabs/encodingjson-generator/#0
type Shop struct {
	Name string `json:"name"`
	OpenedAt time.Time `json:"opendAt"`
}

// 拡張関数を作成しそれを呼び出すパターン
func (s *Shop) MarshalJson() ([]byte, error) {
	// subnameを追加
	return json.Marshal(&struct {
		Name string `json:"name"`
		SubName string `json:"subname"`
		OpendAt time.Time `json:"opendAt"`
	} {
		s.Name, "sub" + s.Name, s.OpenedAt,
	})
}

func sampleCustom() {
	shop := Shop{
		"abc", time.Now(),
	}

	// {"name":"abc","opendAt":"2021-07-26T09:21:26.157004+09:00"}
	b, _ := json.Marshal(&shop)
	log.Println(string(b))

	// {"name":"abc","opendAt":"2021-07-26T09:21:26.157004+09:00"}
	b, _ = shop.MarshalJson()
	log.Println(string(b))

}