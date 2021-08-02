# gui-fyne

- commands
  - create package: `fyne package -os windows -icon cycle.png`
    - 実行ファイルを作成

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
