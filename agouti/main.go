package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sclevine/agouti"
)

var CHROME_OPTION_ARG_HEADRESS = []string{
	"--headless",
	"--disable-gpu",
}

func main() {
	sample(false)
}

func sample(isHeadless bool) {
	option_args := []string{}
	if isHeadless {
		option_args = append(option_args, CHROME_OPTION_ARG_HEADRESS...)
	}

	options := agouti.ChromeOptions(
		"args", option_args,
	)

	driver := agouti.ChromeDriver(options)
	defer driver.Stop()

	if err := driver.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	page, err := driver.NewPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	page.Navigate("https://www.google.com/?hl=ja")

	// スクリプトを実行
	// 引数のコンバートについて：
	//   - $を付ける必要はない。明確にするためサンプルでつけているだけ
	//   - alert('$word') にすると置換が発生しない。convertはjavascript内の引数に値をセットする処理で、左記では文字列扱いのため
	var result interface{}
	if err := page.RunScript(
		"alert($word);",
		map[string]interface{}{"$word": "hello world"},
		&result); err != nil {
		log.Println(err)
	} else {
		log.Println(result)
	}
	time.Sleep(time.Second * 5)

	// タイトルを取得
	fmt.Println(page.Title())
}
