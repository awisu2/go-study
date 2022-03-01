# go-study/json

- [json - The Go Programming Language](https://golang.org/pkg/encoding/json/)

**NOTE**

- decode: `err := json.Marshal([]byte, interface)`
- encode: `jByte, err := json.Marshal([]byte)`

## libraries

replaceable and more quick

- [json\-iterator/go: A high\-performance 100% compatible drop\-in replacement of "encoding/json"](https://github.com/json-iterator/go)
- [goccy/go\-json: Fast JSON encoder/decoder compatible with encoding/json for Go](https://github.com/goccy/go-json)

## samples

interface が struct だった場合以下のように json 変換時のフィールド名を記載することができる

```go
type Message struct {
	Id int `json:"id"`
	Word string `json:"word"`
	Suffix string `json:"suffix"`
}
```

note: optional なので記載しなくてもよいが、public アクセスのための先頭大文字もそのまま反映
