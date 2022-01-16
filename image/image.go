package image

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

// 保存オプション
type SaveOption struct {
	Path    string
	Format  Format
	Quality int // quality 1~100 (jpeg用)
}

type CreateOption struct {
	X0 int
	Y0 int
	X1 int
	Y1 int
	SaveOption
}

// 画像作成
func Create(option *CreateOption) error {
	rect := image.Rect(option.X0, option.Y0, option.X1, option.Y1)
	img := image.NewRGBA(rect)
	img.SetRGBA(1, 1, color.RGBA{255, 0, 0, 0})

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
