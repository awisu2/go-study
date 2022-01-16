# go-study

自分用 go 勉強リポジトリです

## docs

### 基本動作

- [array-slice](./array-slice): array/slice
- [goroutine](./goroutine): goroutine
- [interface](./interface): interface
- [images](./images): imags
- [json](./json): json への変換及びデコード時処理
- [pointer-behavior](./pointer-behavior): 各種値の参照の取り扱い
- [regexp](./regexp): 正規表現
- [sharedGo](./sharedGo): node の native モジュールで動作させることを目標とした実験
- [struct](./struct): struct 構造体
- [time](./time): time
- [test](./test): test
- [tour-of-go](./tour-of-go): tour-of-go 用
- [url](./url): url の取り扱い

### パッケージ動作

- [agouti](./agouti): ブラウザコントロール API サポート(テストや自動挙動確認用)
- [argparse](./argparse): コマンド実行時の引数解析 python に準拠
- [cobra](./cobra): コマンドライン構築サポート (階層的コマンド、help の自動補完、及び viper との連携などが強力)
- [convert](./convert): 型変換実験場
- [echo](./echo): web フレームワーク
- [echo-jwt](./echo-jwt): echo で jwt 認証
- [echo-middleware](./echo-middleware): echo の middleware
- [getstart](./getstart): チュートリアル通りやってみた
- [gorm](./gorm): ORM ライブラリ
- [gorm-simple](./gorm-simple): gorm のかんたん実装
- [gui-fyne](./gui-fyne): go の gui ライブラリ fyne
- [lorca](./lorca): go の gui ライブラリ lorca
- [server-echo](./server-echo): echo を実際利用するため色々構築
- [viper](./viper): コンフィグ値の取り扱い(環境変数や特定の config ファイルからの取得など, cobra との連携も可能)

## インストールなど基本事項

- install: [Download and install - The Go Programming Language](https://golang.org/doc/install)
- commands: 詳細は後述
  - create mod: `go mod init {domain/module}`
  - update packages: `go mod tidy`
  - run mod: `go mod .`
- packages: [pkg.go.dev](https://pkg.go.dev/)
- 環境変数(GOPATH, GOROOT, GOBIN...): 後述
- moduleGO PATH GOPATH に縛られたくない: 無理。(どうしても必要なら GOPATH の自動切り替え環境を用意)
- test: `go test`
  - 命名規則: file: `Xxx_test.go`, func: `TestXxx`
    - ファイル名は何でも良い、func 名が package 内でかぶると (普通に) redeclared error
  - Package がずれていると動作しないみたい: `found packages main (simple_sample.go)...`
    - simple_sample_test.go の　 package を main に変えたら動いた
- vscode での補完がおかしい: vscode workspace root = go pacakgae root である必要があるとのこと
  - このリポジトリの用に小ディレクトリに package 置いていると発生する
  - add folder to workspace で root になるようにしてやれば OK

## links

- [The Go Programming Language](https://golang.org/)
- [A Tour of Go](https://go-tour-jp.appspot.com/welcome/1)

## commands

go には module モードと、GOPATH モードがある。デフォルトでは module モードで、GOPATH モードは GOPATH 配下にすべてを詰め込む旧仕様。
module モードにすることで、module のバージョンなどを分離できる。。。らしい。 go.mod がその役目？

- init: `go mod init {domain/module}`
- run: `go run .`
- tidy mod: `go mod tidy`
  - after writed `imoprt` in code
- edit mod: `go mod edit {command}`
  - replace module path: `go mod edit -replace example.com/greetings=../greetings`
- build: `go build`
  - check install path (it builded): `go list -f {{.Target}}`
    - これを環境変数の PATH に入れておけばビルドしたコマンドが利用できる

## install

[Download and install - The Go Programming Language](https://golang.org/doc/install)

### linux

```bash
# install
GOFILE="go1.16.5.linux-amd64.tar.gz"
cd
wget https://golang.org/dl/${GOFILE}
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf ${GOFILE}

# add path (if not exists)
cat - << 'EOF' > ~/.bash_profile
# go
export PATH=$PATH:/usr/local/go/bin
EOF

# after restart
go version
```

## 環境変数

- GOPATH: go が path 解決する際に root として扱うディレクトリ(ワークスペース)
  - 配下に package やバイナリ、開発コードが存在することが前提として動作する
  - 利用可能なディレクトリであればどこでも OK
  - os(process)ごとに１つしか指定できない
    - (開発元の異なるなど)ディレクトリを分けたい場合逐一切り替える必要あり
- GOROOT: よくわからないが、go 本体のインストールパスが割り当てられていた
  - GOROOT の設定は不要との記事は散見する。実際しなくて問題ない
- GOBIN: built 時の install 先 (指定しなければ `${GOPATH}/bin` になるっぽい)
