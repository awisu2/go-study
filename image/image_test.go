package image

import (
	"testing"
)

func TestCreate(t *testing.T) {
	option := CreateOption{
		X0: 0,
		Y0: 0,
		X1: 100,
		Y1: 50,
		SaveOption: SaveOption{
			Path:    "sample.jpg",
			Format:  Jpg,
			Quality: 70,
		},
	}
	err := Create(&option)
	if err != nil {
		t.Errorf("%v", err)
	}
}
