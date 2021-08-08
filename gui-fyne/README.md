# gui-fyne

```go
package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// appの作成(1threadで1つ、contentを作成する前提だったりする)
	myApp := app.New()

	// windowの中身の作成
	myWindow := myApp.NewWindow("Hello")

	// windowへ中身をセット
	myWindow.SetContent(widget.NewLabel("Hello"))

	// 表示
	myWindow.Show()
	myApp.Run()

	// 終了
	fmt.Println("window finish")
}
```

- commands
  - create package: `fyne package -os windows -icon cycle.png`
    - 実行ファイルを作成: 
  - 画像や音声ファイルなどをbundling(goコード化): `fyne bundle some.png > resources.go`
    - [Bundling resources | Develop using Fyne](https://developer.fyne.io/tutorial/bundle)
	- なにがメリット？: goのアプリケーション最終的には基本一つの実行ファイルになる。fyneもそうしたい。
	  - しかしアイコンなどコード外のファイルが存在すると実行ファイルとは別のファイルがひつようになる。それらをbundlingすることで、アプリの中に含めることが可能になる。
    - "resourceSomePng" という変数が作成され、`img := canvas.NewImageFromResource(resourceSomePng)` というように利用できる
	- 同じファイルに追加で出力する際はappendをつける: `fyne bundle -append image2.png >> resources.go`
	  - 複数管理する場合は、shellスクリプトなどで自動化することをおすすめ
- 画像: 
  - FillModeを設定しないと潰れる: `img.FillMode = canvas.ImageFillOriginal`

## links

- [Fyne](https://fyne.io/)
- [Fyne toolkit documentation for developers | Develop using Fyne](https://developer.fyne.io/)

## insatll

- fyne コマンドのインストール: `go get fyne.io/fyne/v2/cmd/fyne`

## getting start

main.go

```go
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :")
		}),
	))

	w.ShowAndRun()
}
```
