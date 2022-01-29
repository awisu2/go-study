package agouti

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestGetTitle(t *testing.T) {
	userData, _ := filepath.Abs("./tmp/userData")
	title, err := sampleGetTitle("https://www.google.com/?hl=ja", &DriverOption{Height: 400, Width: 100, UserDataDir: userData})
	if err != nil {
		t.Errorf("any error happen. %v", err)
	}

	fmt.Println(title)

	if title == "" {
		t.Errorf("title not exists")
	}

}
