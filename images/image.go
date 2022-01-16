package image

import (
	"image"
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
}

func CreateImage(size *Size) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, size.Width, size.Height))
	return img
}

// 画像作成
func Create(option *CreateOption) error {
	img := CreateImage(&option.Size)

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

func Resize(img image.Image, size *Size) (*image.RGBA, error) {
	resizedImg := CreateImage(size)
	draw.CatmullRom.Scale(resizedImg, resizedImg.Bounds(), img, img.Bounds(), draw.Over, nil)
	return resizedImg, nil
}
