package main

// 自動補完に任せるとちょいちょいv2が抜けてコンパイルエラーになるので気をつける
import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"os"
	"time"

	"fyne.io/fyne/theme"
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
	// showSecond(app)
	// showWidget(app)
	// showCanvas(app)
	// showLayouts(app)
	showTutorials(app)

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

func showLayouts(app fyne.App) {
	config := Config{
		Title:      "layouts",
		WindowSize: fyne.NewSize(300, 300),
	}

	text := canvas.NewText("1", color.Black)
	text2 := canvas.NewText("2", color.Black)

	// 横並び horizontal box layout
	// hBox := layout.NewHBoxLayout()

	// 縦並び vertical box layout
	vBox := layout.NewVBoxLayout()
	spacer := layout.NewSpacer()

	// grid
	// container化することで、他のlayoutに混ぜることが可能
	// 引数は横のカラム数で等間隔にっセットされる
	grid := container.New(
		layout.NewGridLayout(3),
		canvas.NewText("g1", color.Black),
		canvas.NewText("g2", color.Black),
		canvas.NewText("g3", color.Black),
		canvas.NewText("g4", color.Black),
		canvas.NewText("g5", color.Black),
		canvas.NewText("g6", color.Black),
		canvas.NewText("g7", color.Black),
		canvas.NewText("g8", color.Black),
	)

	// grid wrap 1ブロックあたりのサイズを指定しそれが収まるようにgrid表示
	gridWrap := container.New(
		layout.NewGridWrapLayout(fyne.NewSize(50, 50)),
		canvas.NewText("g1", color.Black),
		canvas.NewText("g2", color.Black),
		canvas.NewText("g3", color.Black),
		canvas.NewText("g4", color.Black),
		canvas.NewText("g5", color.Black),
		canvas.NewText("g6", color.Black),
		canvas.NewText("g7", color.Black),
		canvas.NewText("g8", color.Black),
	)

	// border: midを中心に上下左右にcontentを配置することが可能
	//
	// Layout宣言時に、配置するContentObjectを渡す必要がある(サイズ取得のため？)
	// > 別途、表示するコンテンツは、Container.Newの引数にも渡す必要あり
	// nilを渡した箇所は、何も表示されない
	// style と違い, 上下左右の順なので注意
	mid := canvas.NewText("b_mid", color.Black)
	top := canvas.NewText("b_top", color.Black)
	bottom := canvas.NewText("b_bottom", color.Black)
	left := canvas.NewText("b_left", color.Black)
	right := canvas.NewText("b_right", color.Black)
	border := container.New(layout.NewBorderLayout(top, bottom, left, right),
		top, bottom, left, right, mid)

	// form: いわゆる左に名前、右に入力フォームを設定するとき用のlayout
	//
	// 挙動はNewGridLayout(2) に似ているが、カラム幅は左のみ右のみの最大で設定され、高さもそれぞれの行の最大で設定される
	//
	form := container.New(layout.NewFormLayout(),
		canvas.NewText("Label 1", color.Black), canvas.NewText("Value", color.Black),
		canvas.NewText("Label 2", color.Black), canvas.NewText("Value2", color.Black))

	// center: 中央揃え
	// TODO: 高さが設定できない
	centerText := canvas.NewText("center 1", color.Black)
	center := container.New(layout.NewCenterLayout(), centerText)

	// max: 内部要素をすべてcontainerと同じ高さにする
	//
	// 複数要素を指定した場合は上に重なっていく
	max := container.New(layout.NewMaxLayout(),
		canvas.NewImageFromResource(theme.FyneLogo()),
		canvas.NewText("Overlay", color.Black))

	Show(app, container.New(vBox, text, spacer, grid, gridWrap, border, form, center, max, text2), &config)
}

func showTutorials(app fyne.App) {
	config := Config{
		Title:      "tutorials",
		WindowSize: fyne.NewSize(300, 300),
	}

	// widghetの拡張
	// クリック可能なicon, Tapped, TappedSecondary interface{}を追加することで機能追加
	icon := newTappableIcon(theme.FyneLogo())

	img := canvas.NewImageFromResource(resourceCyclePng)
	img.FillMode = canvas.ImageFillOriginal

	diagonal := container.New(
		&diagonal{},
		widget.NewLabel("top left"),
		widget.NewLabel("middle"),
		widget.NewLabel("bottom right"),
	)

	Show(app, container.New(layout.NewVBoxLayout(), icon, diagonal, img), &config)
}

type diagonal struct {
}

// 最初に表示領域としてMiniSizeを設定
func (d *diagonal) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}

// 実際の配置
func (d *diagonal) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, containerSize.Height-d.MinSize(objects).Height)
	for _, o := range objects {
		size := o.MinSize()
		o.Resize(size)
		o.Move(pos)

		// 追加したcontentのサイズ分加算(右下に連鎖していく感じ)
		pos = pos.Add(fyne.NewPos(size.Width, size.Height))
	}
}
