# test

[testing \- Go 言語](https://xn--go-hh0g6u.com/pkg/testing/)

```bash
go test
```

- ベンチマーク: オプション指定することでベンチマークを追加で実行可能
- 命名ルール
  - テスト関数は _Test_ で始まること
  - テスト関数を含むファイルは _\_test.go_ で終わること
  - ベンチマーク:
    - 関数は _Benchmark_ で始まること
    - ファイル名は test と同じ
  - 命名ルールに則ったファイルは go test のときにのみビルド対象となる
- テスト引数
  - testing.T: 通常のテスト用
  - testing.B: ベンチマークを実行し処理時間を図る

## テストコマンド

- 全テスト: `go test`
- パッケージや関数名を指定してテスト: `go test {package} -run {regexp}`
  - 正規表現は部分一致,大/小文字区別あり,関数名に一致 (何も指定しない "" 場合は全てに一致)
  - パッケージは `./` とすることでカレントディレクトリ指定になる
  - パッケージをディレクトリではなく指定すると `go get` で取得してのテストを実行する
- ベンチマーク: `go test --bench {regexp}`
- パッケージチートシート: 正規表現ではないっぽいが ".", "...", "\*" などがメタ文字として使える
  - 同値表現: `go test ./` == `go test ./.` =？`go test./*go`
  - 自分を含め直下すべて: `go test ./...`
    - サブディレクトリ以降すべて: `go test ./**/*` (... があるのであまり意味がない)
  - 特定のサブパッケージ: `go test ./sub`
  - 特定のファイル: `go test ./any*`
- regexp チートシート(--bench も-run も同条件)
  - 全テスト/ベンチマーク: `go test ./... --bench .`
  - ベンチマークのみ: `go test ./... -run ^$ --bench .` (前方一致と後方一致で挟み込んで対象なしに)
- オプションドキュメント: [go command \- cmd/go \- pkg\.go\.dev](https://pkg.go.dev/cmd/go#hdr-Testing_flags)

## サンプルコード

```go
func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}

func BenchmarkAbs(b *testing.B) {
    got := Abs(-1)
    if got != 1 {
        b.Errorf("Abs(-1) = %d; want 1", got)
    }
}
```
