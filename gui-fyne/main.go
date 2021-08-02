package main

// 自動補完に任せるとちょいちょいv2が抜けてコンパイルエラーになるので気をつける
import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	Title      string
	WindowSize fyne.Size
}

func (config *Config) CreateApp() fyne.App {
	return app.New()

}

func (config *Config) CreateWindow(app fyne.App) fyne.Window {
	w := app.NewWindow(config.Title)
	w.Resize(config.WindowSize)
	return w
}

func CreateConfig(config *Config) *Config {
	if config.WindowSize.IsZero() {
		config.WindowSize.Height = 1000
		config.WindowSize.Width = 1000
	}
	return config
}

func main() {
	config := CreateConfig(&Config{
		Title: "base",
	})

	// appが作成されていないと、特定のコンテンツ作成時エラーになる
	// Fyne error:  Attempt to access current Fyne app when none is started
	app := config.CreateApp()

	// サブ window を表示
	showSecond(app)
	showWidget(app)
	showCanvas(app)

	// メインwindowを表示
	content := contentBoxLayout()
	ShowAndRun(app, content, config)
}

func ShowAndRun(app fyne.App, content fyne.CanvasObject, config *Config) {
	w := config.CreateWindow(app)
	w.SetContent(content)

	// show and run
	// w.Show()
	// app.Run()
	w.ShowAndRun()
}

// show だけ実行することでふくすwindowを表示できる
func Show(app fyne.App, content fyne.CanvasObject, config *Config) fyne.Window {
	w := config.CreateWindow(app)
	w.SetContent(content)
	go w.Show()
	return w
}

func contentIntroduction() fyne.CanvasObject {
	return widget.NewLabel("Hello Fyne!")
}

func contentBoxLayout() fyne.CanvasObject {
	// 普通のテキスト
	// canvasは表示領域,swift-uiのviewに近い
	text1 := canvas.NewText("Hello", color.Black)
	text2 := canvas.NewText("There", color.Black)
	text3 := canvas.NewText("(right)", color.Black)
	text4 := canvas.NewText("centered", color.Black)

	// HBoxLayout: 横並べ
	// container.New: layout込みでcontainerを作成
	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	// 中央寄せのテキスト作成
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())

	// VBoxLayout: 縦並べ
	return container.New(layout.NewVBoxLayout(), content, centered)
}

func showSecond(app fyne.App) {
	config2 := CreateConfig(&Config{
		Title:      "second",
		WindowSize: fyne.NewSize(300, 200),
	})

	content2 := contentIntroduction()
	w := Show(app, content2, config2)
	go ChangeContent(w)
}

func ChangeContent(window fyne.Window) {
	time.Sleep(time.Second)
	window.SetContent(widget.NewLabel("1"))

	time.Sleep(time.Second)
	window.SetContent(widget.NewLabel("2"))

	time.Sleep(time.Second)
	window.SetContent(widget.NewLabel("3"))

	time.Sleep(time.Second)
	// NewWithoutLayout でabsoluteみたいになる
	text1 := canvas.NewText("1", color.Black)
	text2 := canvas.NewText("2", color.Black)
	text3 := canvas.NewText("3", color.Black)
	text2.Move(fyne.NewPos(15, 15))
	text3.Move(fyne.NewPos(30, 30))
	window.SetContent(container.NewWithoutLayout(text3, text2, text1))
}

func showWidget(app fyne.App) {
	config := Config{
		Title:      "widget",
		WindowSize: fyne.NewSize(300, 300),
	}

	entry := widget.NewEntry()
	Show(app, container.New(layout.NewVBoxLayout(), entry), &config)
}

func showCanvas(app fyne.App) {
	config := Config{
		Title:      "canvas",
		WindowSize: fyne.NewSize(300, 300),
	}

	rect := canvas.NewRectangle(color.Black)

	text := canvas.NewText("aaa", color.Black)
	// 左: fyne.TextAlignLeading
	// 中央: fyne.TextAlignCenter
	// 右: fyne.TextAlignTrailing
	text.Alignment = fyne.TextAlignTrailing
	text.TextStyle = fyne.TextStyle{Italic: true, Bold: true}

	line := canvas.NewLine(color.Black)
	line.StrokeWidth = 5

	circle := canvas.NewCircle(color.Black)
	circle.StrokeColor = color.Gray{0x99}
	circle.StrokeWidth = 30

	path := "cycle.png"
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	i, _, err := image.Decode(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	img := canvas.NewImageFromFile(path)
	// img2 := canvas.NewImageFromReader(f, t)
	img3 := canvas.NewImageFromImage(i)
	imgLine1 := canvas.NewLine(color.Black)

	img.FillMode = canvas.ImageFillOriginal
	// img2.FillMode = canvas.ImageFillOriginal
	img3.FillMode = canvas.ImageFillOriginal

	// 領域に合わせて可変する画像
	raster := canvas.NewRasterWithPixels(
		func(_, _, w, h int) color.Color {
			return color.RGBA{uint8(rand.Intn(255)),
				uint8(rand.Intn(255)),
				uint8(rand.Intn(255)), 0xff}
		})

	raster2 := canvas.NewRasterFromImage(i)

	// グラデーション？
	gradient := canvas.NewHorizontalGradient(color.Black, color.Transparent)

	Show(app, container.New(layout.NewVBoxLayout(),
		rect,
		text,
		line,
		circle,
		img, imgLine1, img3,
		raster, raster2,
		gradient,
	), &config)
}

func getImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	i, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return i, nil
}
