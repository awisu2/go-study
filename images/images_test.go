package images

import (
	"testing"
)

var createOption = CreateOption{
	Size: Size{
		Width:  100,
		Height: 50,
	},
	SaveOption: SaveOption{
		Path:    "sample.png",
		Format:  Png,
		Quality: 70,
	},
	Color: RgbaGreen,
}

func TestCreate(t *testing.T) {
	err := Create(&createOption)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestOpen(t *testing.T) {
	err := Create(&createOption)
	if err != nil {
		t.Errorf("%v", err)
	}

	_, format, err := Open(createOption.Path)
	if err != nil {
		t.Errorf("%v", err)
	}
	if format != "png" {
		t.Errorf("got %v want %s", format, "png")
	}
}

func TestResize(t *testing.T) {
	// テスト用画像作成
	err := Create(&createOption)
	if err != nil {
		t.Errorf("%v", err)
	}
	img, _, err := Open(createOption.Path)
	if err != nil {
		t.Errorf("%v", err)
	}

	// リサイズ
	size := Size{
		Width:  50,
		Height: 50,
	}
	resizedImg, err := Resize(img, &size)
	if err != nil {
		t.Errorf("%v", err)
	}

	// サイズ確認
	bounds := resizedImg.Bounds()
	if bounds.Dx() != size.Width || bounds.Dy() != size.Height {
		t.Error("resize missing")
	}

	// リサイズ済み画像保存
	err = Save(resizedImg, &SaveOption{
		Path:   "sample2.png",
		Format: Png,
	})
	if err != nil {
		t.Errorf("save error: %v", err)
	}
}
