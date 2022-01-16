package images

import "image/color"

// color.RGBAは全パラメータ uint8
//   0-255 で扱われ、Aは 0 の時透明
var (
	RgbaRed = color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}
	RgbaGreen = color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 255,
	}
	RgbaBlue = color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 255,
	}
)
