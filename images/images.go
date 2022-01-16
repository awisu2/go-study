package images

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

// 保存オプション
type SaveOption struct {
	Path    string
	Format  Format
	Quality int // quality 1~100 (jpeg用)
}

type CreateOption struct {
	Size
	SaveOption
	Color color.Color
}

func CreateImage(size *Size) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, size.Width, size.Height))
	return img
}

// 画像作成
func Create(option *CreateOption) error {
	img := CreateImage(&option.Size)
	SetRgbasFull(img, option.Color)

	err := Save(img, &option.SaveOption)
	if err != nil {
		return err
	}

	return nil
}

// ファイルに保存
func Save(img image.Image, option *SaveOption) (err error) {
	f, err := os.Create(option.Path)
	if err != nil {
		return err
	}
	defer f.Close()

	if option.Format == Jpg {
		err = jpeg.Encode(f, img, &jpeg.Options{Quality: option.Quality})
	} else if option.Format == Png {
		err = png.Encode(f, img)
	}
	if err != nil {
		return err
	}

	return nil
}

// 特定のファイルを読み込み
func Open(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	image, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return image, nil
}

// 画像のリサイズ
func Resize(img image.Image, size *Size) (*image.RGBA, error) {
	// 変換後画像オブジェクト作成
	resizedImg := CreateImage(size)
	bounds := resizedImg.Bounds()

	// もと画像を拡大/縮小して書き込み
	// todo: Kernelを切り替えられるようにしても良い
	// 色が先に塗られていた場合dstが優先される
	// 第5引数の draw.Op: alphaが効く形式のときに有効
	//   draw.Over: 上書き(色は交じる)
	//   draw.Src: dstが優先され色の合成はされない
	draw.CatmullRom.Scale(resizedImg, bounds, img, img.Bounds(), draw.Src, nil)
	return resizedImg, nil
}

// 一定範囲に色をセット
func SetRgbas(img *image.RGBA, dr image.Rectangle, color color.Color) {
	for y := dr.Min.Y; y < dr.Max.Y; y++ {
		for x := dr.Min.X; x < dr.Max.X; x++ {
			img.Set(x, y, color)
		}
	}
}

// 全範囲に色をセット
func SetRgbasFull(img *image.RGBA, color color.Color) {
	SetRgbas(img, img.Bounds(), color)
}
