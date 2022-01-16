# test

[testing \- Go 言語](https://xn--go-hh0g6u.com/pkg/testing/)

```bash
go test
```

- テスト用のファイルは通常はビルド対象とならず test 時のみビルドされる
- ベンチマーク: オプション指定することでベンチマークを追加で実行可能
- 命名ルール
  - テスト関数は _Test_ で始まること
  - テスト関数を含むファイルは _\_test.go_ で終わること
  - ベンチマーク:
    - 関数は _Benchmark_ で始まること
    - ファイル名は test と同じで追加する感じ
- test ファイル のパッケージはテスト対象と同じにして関数呼び出しを簡易にするのが基本と思われる

## テストコマンド

[go command \- cmd/go \- pkg\.go\.dev](https://pkg.go.dev/cmd/go#hdr-Testing_flags)

- ヘルプ: `go help test` ただあまり情報が多くない(--bench すら出てこない)
- 全テスト: `go test`
- パッケージや関数名を指定してテスト: `go test {package} -run {regexp}`
  - 正規表現は部分一致,大/小文字区別あり,関数名に一致 (何も指定しない "" 場合は全てに一致)
  - パッケージは `./` とすることでカレントディレクトリ指定になる
  - パッケージをディレクトリではなく指定すると `go get` で取得してのテストを実行する
- ベンチマーク: `go test --bench {regexp}`
- パッケージチートシート: 正規表現ではないっぽいが ".", "...", "\*" などがメタ文字として使える
  - 同値表現: `go test ./` == `go test ./.` =？`go test./*go`
  - 自分を含め直下すべて: `go test ./...` (`go test ...` とすると世の中に存在する全てのパッケージが対象となるので気をつける)
    - サブディレクトリ以降すべて: `go test ./**/*` (... があるのであまり意味がない)
  - 特定のサブパッケージ: `go test ./sub`
  - 特定のファイル: `go test ./any*`
- regexp チートシート(--bench も-run も同条件)
  - 全テスト/ベンチマーク: `go test ./... --bench .`
  - ベンチマークのみ: `go test ./... -run ^$ --bench .` (前方一致と後方一致で挟み込んで対象なしに)

### オプション

- 実行されるテスト個々の情報を出力: `-v`
- 指定回数実行: `-count n`
- 利用 cpu 数(GOMAXPROCS): `-cpu 1,2,4`

### 追加出力

- `-benchmem`: --bench でベンチマークが実行された時、メモリ割当情報が追加される

## サンプルコード

_main_test.go_

```go
package main

import (
	"testing"
)

// テスト
func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}

// ベンチマーク
func BenchmarkAbs(b *testing.B) {
    got := Abs(-1)
    if got != 1 {
        b.Errorf("Abs(-1) = %d; want 1", got)
    }
}
```

## todo

- [] 追加出力の詳細をまとめる
