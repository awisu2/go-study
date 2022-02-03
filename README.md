# go-study

自分用 go 勉強リポジトリです

## docs

- [basics](./docs/basics.md): インストールや基本事項
- [module](./docs/module.md): module の取り扱いについて(get/install コマンドなど)
- [特殊ディレクトリ](./docs/specialDirectory.md): 特殊ディレクトリの取得

## contents

### 基本動作

- [array-slice](./array-slice): array/slice
- [goroutine](./goroutine): goroutine
- [helloworld](./helloworld): hello world !
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

### 各種 module を使ってみる

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
