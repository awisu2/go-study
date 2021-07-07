# go-study

自分用 go 勉強リポジトリです

- install: [Download and install - The Go Programming Language](https://golang.org/doc/install)
- commands: 後述
- packages: [pkg.go.dev](https://pkg.go.dev/)
- GOPATH, GOROOT
  - GOPATH: go が path 解決する際に root として扱うディレクトリ(ワークスペース)
    - 配下に package やバイナリ、開発コードが存在することが前提として動作する
    - 利用可能なディレクトリであればどこでも OK
    - os(process)ごとに１つしか指定できない
      - (開発元の異なるなど)ディレクトリを分けたい場合逐一切り替える必要あり

簡単サンプルコード

```go
func main() {
  msg := "foo"
  fmt.Println(msg)


}

func init() {
  log.SetPrefix("sample")
  log.SetFlags(0)
}
```

## commands

- init: `go mod init {domain/module}`
- run: `go run .`
- tidy mod: `go mod tidy`
  - before write `imoprt` in code
- edit mod: `go mod edit {command}`
  - replace module path: `go mod edit -replace example.com/greetings=../greetings`
