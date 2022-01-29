package goquery

import (
	"testing"
)

func TestLoadHtml(t *testing.T) {
	_, err := LoadDocument("./index.html")
	if err != nil {
		t.Errorf("load missing. %v", err)
	}
}

func TestGetHtml(t *testing.T) {
	_, err := GetDocument("https://www.google.com/?hl=ja")
	if err != nil {
		t.Errorf("load missing. %v", err)
	}
}

func TestSampleEach(t *testing.T) {
	liTexts, liBtexts, err := SampleEach()
	if err != nil {
		t.Errorf("Fatal error. %v", err)
	}

	if len(liTexts) != 8 || len(liBtexts) != 8 {
		t.Errorf("selector's length is fail")
	}
}
