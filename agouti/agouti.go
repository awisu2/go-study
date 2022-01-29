package agouti

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
)

type Driver string

const (
	DRIVER_CHROME Driver = "chrome"
	DRIVER_OTHER  Driver = "other"
)

type DriverOption struct {
	Headless bool // ヘッドメニューを非表示(max/linuxのみ)
	Width    int  // 幅
	Height   int  // 高さ
}

func (opt *DriverOption) FixDefault() {
}

// chrome driverを起動時、5秒待機その後、タイトルを返却
func getTitle(url string, opt *DriverOption) (string, error) {
	// ドライバを作成し、起動
	driver := createChromeDriver(DRIVER_CHROME, opt)
	if err := driver.Start(); err != nil {
		return "", err
	}
	defer driver.Stop()

	// 新規ページ(タブ/ウィンド)を起動
	page, err := driver.NewPage()
	if err != nil {
		return "", err
	}

	// 特定のサイトに移動(指定しない場合はブランクページ)
	page.Navigate(url)

	// スクリプトを実行
	var result interface{}
	if err := runAlert(page, "hello world", &result); err != nil {
		log.Panic(err)
	}
	log.Println(result)

	time.Sleep(time.Second * 5)

	// HTMlを取得
	html, err := page.HTML()
	if err != nil {
		return "", err
	}
	// htmlを解析してタイトル取得
	title1, err := getTitlefromHtml(html)
	if err != nil {
		return "", err
	}

	// タイトルを取得
	title2, err := page.Title()
	if err != nil {
		return "", err
	}

	if title1 != title2 {
		return "", fmt.Errorf("not equal title. %v, %v", title1, title2)
	}

	return title1, nil
}

func createChromeDriver(driver Driver, opt *DriverOption) *agouti.WebDriver {
	// 起動時フラグ
	args := []string{}
	if opt.Headless {
		args = append(args, "--headless", "--disable-gpu")
	}
	if opt.Width > 0 && opt.Height > 0 {
		args = append(args, fmt.Sprintf("--window-size=%d,%d", opt.Width, opt.Height))
	}

	log.Printf("%v\n", args)
	options := agouti.ChromeOptions(
		"args", args,
	)

	// ドライバを起動
	return agouti.ChromeDriver(options)
}

// スクリプトの実行
//
// どちらもpageにスクリプトと変数を割り当てるだけで、動作はjavascript
// よって、alert('$word') にすると文字列扱いになりalertは "$word" が出力される
func runAlert(page *agouti.Page, word string, result *interface{}) error {
	body := "alert($word);"
	arguments := map[string]interface{}{"$word": word}
	return page.RunScript(body, arguments, result)
}

// qoqueryにより、htmlを解析してtitle取得
func getTitlefromHtml(html string) (string, error) {
	// documentを取得
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}

	selection := doc.Find("title")

	// 複数要素を取れますよという参考に(Eachコマンドなどでループ可能)
	if selection.Length() != 1 {
		return "", fmt.Errorf("not only one title!")
	}

	return selection.Text(), nil
}
