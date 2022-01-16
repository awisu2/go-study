package image

type Size struct {
	Width  int
	Height int
}

// todo: constで保管しておく方法
var (
	Hd = Size{
		Width:  2560,
		Height: 1440,
	}
	FullHd = Size{
		Width:  1920,
		Height: 1000,
	}
	K2 = Size{
		Width:  2560,
		Height: 1440,
	}
	K4 = Size{
		Width:  3840,
		Height: 2160,
	}
	K8 = Size{
		Width:  7680,
		Height: 4320,
	}
)
