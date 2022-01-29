package agouti

import (
	"fmt"
	"testing"
)

func TestGetTitle(t *testing.T) {
	title, err := getTitle("https://www.google.com/?hl=ja", &DriverOption{Height: 600, Width: 300})
	if err != nil {
		t.Errorf("any error happen. %v", err)
	}

	fmt.Println(title)

	if title == "" {
		t.Errorf("title not exists")
	}

}
