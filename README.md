# go-study

自分用 go 勉強リポジトリです

- install: [Download and install - The Go Programming Language](https://golang.org/doc/install)
- commands: 後述
- packages: [pkg.go.dev](https://pkg.go.dev/)
- 環境変数(GOPATH, GOROOT, GOBIN...): 後述
- GOPATH に縛られたくない: 無理。(どうしても必要なら GOPATH の自動切り替え環境を用意)
- test: `go test`
  - 命名規則: file: `Xxx_test.go`, func: `TestXxx`
    - ファイル名は何でも良い、func 名が package 内でかぶると (普通に) redeclared error
  - Package がずれていると動作しないみたい: `found packages main (simple_sample.go)...`
    - simple_sample_test.go の　 package を main に変えたら動いた

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
- build: `go build`
  - check install path (it builded): `go list -f {{.Target}}`
    - これを環境変数の PATH に入れておけばビルドしたコマンドが利用できる

## 環境変数

- GOPATH: go が path 解決する際に root として扱うディレクトリ(ワークスペース)
  - 配下に package やバイナリ、開発コードが存在することが前提として動作する
  - 利用可能なディレクトリであればどこでも OK
  - os(process)ごとに１つしか指定できない
    - (開発元の異なるなど)ディレクトリを分けたい場合逐一切り替える必要あり
- GOROOT: よくわからないが、go 本体のインストールパスが割り当てられていた
  - GOROOT の設定は不要との記事は散見する。実際しなくて問題ない
- GOBIN: built 時の install 先 (指定しなければ `${GOPATH}/bin` になるっぽい)
