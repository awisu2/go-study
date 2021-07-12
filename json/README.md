# go-study/json

- 基本の 2 メソッド
  - デコード: `err := json.Marshal([]byte, interface)`
  - エンコード: `jByte, err := json.Marshal([]byte)`

interface が struct だった場合以下のように json 変換時のフィールド名を記載することができる

```go
type Message struct {
	Id int `json:"id"`
	Word string `json:"word"`
	Suffix string `json:"suffix"`
}
```

note: optional なので記載しなくてもよいが、public アクセスのための先頭大文字もそのまま反映

## links

- [json - The Go Programming Language](https://golang.org/pkg/encoding/json/)
