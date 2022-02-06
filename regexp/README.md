# regexp

go での正規表現(regular expression)について

[regexp package \- regexp \- pkg\.go\.dev](https://pkg.go.dev/regexp)

## sample

先に pattern を指定したオブジェクトを作成する or pattern と同時に実行する のどちらかに分かれる

```go
var re = regexp.MustCompile(`foo.*`)
matched, err := re.Match([]byte(`seafood`))
// or
matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
```

## methods

### パターンオブジェクト作成

- Compile: expr を解析して Regexp オブジェクトを返却 `func Compile(expr string) (*Regexp, error)`
- CompilePOSIX: Compile とほぼ同様だが、POSIX ERE (egrep)構文で扱う.`func CompilePOSIX(expr string) (*Regexp, error)`
- MustCompile: Compile とほぼ同様だが expr が間違っている場合は panic
- MustCompilePOSIX: CompilePOSIX とほぼ同様だが expr が間違っている場合は panic

### サポート

- QuoteMeta: 正規表現を escape する `func QuoteMeta(s string) string`

### 正規表現処理

- Match: `func Match(pattern string, b []byte) (matched bool, err error)`
- MtchString: `func MatchString(pattern string, s string) (matched bool, err error)`
- [Expand](https://pkg.go.dev/regexp#Regexp.Expand): content と match 情報をあわせて template の $xxx の部分を置き換える。その結果を dst に追加して返却
  - match は事前に取得しておく
  - `func (re \*Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte`
