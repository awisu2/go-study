# go-study

my go study

- [The Go Programming Language](https://golang.org/)
- Clean Architecture: [bxcodec/go\-clean\-arch: Go \(Golang\) Clean Architecture based on Reading Uncle Bob's Clean Architecture](https://github.com/bxcodec/go-clean-arch)

## docs

- [A Tour of Go](https://go-tour-jp.appspot.com/welcome/1)
  - official and very good go's first study
- [basics](./docs/basics.md): インストールや基本事項
- [module](./docs/module.md): module の取り扱いについて(get/install コマンドなど)
  - 追記: privte github から go get する方法
- [特殊ディレクトリ](./docs/specialDirectory.md): 特殊ディレクトリの取得
- [rough code](./rough/rough.go): basic sample code

## basic

- [array-slice](./array-slice): array/slice
- [callback](./callback-study/): callback functions
- [commands](./commands/): go commands like `go run .`, `go build .`
- [helloworld](./helloworld): hello world !
- [file](./file): file
- [godoc](./godoc-study/): `go doc` command
- [interface](./interface): interface
- [images](./images): imags
- [json](./json): json への変換及びデコード時処理
- [os](./os-study/): (ex: catch stdout)
- [pointer-behavior](./pointer-behavior): 各種値の参照の取り扱い
- [path](./path): path の取り扱い
- [regexp](./regexp): 正規表現
- [runtime](./runtime-study/): rutime. (ex: what os, architecture running.)
- [reflection](./reflect-study/): for dynamic type
- [sharedGo](./sharedGo): node の native モジュールで動作させることを目標とした実験
- [struct](./struct): struct 構造体
- [string](./string): string の取り扱い
- [switch](./switch-study): switch
- [time](./time): time
- [tour-of-go](./tour-of-go): tour-of-go 用
- [url](./url): url の取り扱い

## test

- [test](./test): test
- [testify](./testify-study/README.md): The testify is supporting test more simply, and add mock and suite.

## groutine, context, errgroup

- [goroutine](./goroutine): goroutine
- [context](./context-study/): context it's can Canceling, Set timeout, and carrie values.
- errgroup
  - [errgroup package \- golang\.org/x/sync/errgroup \- pkg\.go\.dev](https://pkg.go.dev/golang.org/x/sync/errgroup#ex-Group--Pipeline)

## command line

- [argparse](./argparse): コマンド実行時の引数解析 python に準拠
- [cobra](./cobra): コマンドライン構築サポート (階層的コマンド、help の自動補完、及び viper との連携などが強力)
- [viper](./viper): コンフィグ値の取り扱い(環境変数や特定の config ファイルからの取得など, cobra との連携も可能)

## web framework

- [echo](./echo): web フレームワーク
- [echo-jwt](./echo-jwt): echo で jwt 認証
- [echo-middleware](./echo-middleware): echo の middleware
- [server-echo](./server-echo): echo を実際利用するため色々構築
- gin
- aero
- revel
- beego
- goji
- iris

## db

- [gorm](./gorm): ORM ライブラリ
- [gorm-simple](./gorm-simple): gorm のかんたん実装

## gui

- [agouti](./agouti): ブラウザコントロール API サポート(テストや自動挙動確認用)
- [gui-fyne](./gui-fyne): go の gui ライブラリ fyne
- [lorca](./lorca): go の gui ライブラリ lorca

## my try

- [getstart](./getstart): チュートリアル通りやってみた
- [convert](./convert): converting test
