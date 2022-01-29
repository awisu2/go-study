package goquery

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"

	"io/ioutil"
)

// ファイルからDocumentを取得する
func LoadDocument(file string) (*goquery.Document, error) {
	html, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// urlからDocumentを取得
func GetDocument(url string) (*goquery.Document, error) {
	// httpで取得
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status not 200")
	}
	defer res.Body.Close()

	// readerからDocument生成
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func SampleEach() ([]string, []string, error) {
	doc, err := LoadDocument("./index.html")
	if err != nil {
		return nil, nil, err
	}

	// ほぼほぼ js の selector と同等の選択が可能
	liTexts := []string{}
	doc.Find("ul .li").Each(func(_ int, li *goquery.Selection) {
		liTexts = append(liTexts, li.Text())
	})

	// 深い対象でも配列で取得可能
	liBTexts := []string{}
	doc.Find("ul .li b").Each(func(_ int, li *goquery.Selection) {
		liBTexts = append(liBTexts, li.Text())
	})

	return liTexts, liBTexts, nil
}
