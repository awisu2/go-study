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
	Headless    bool   // ヘッドメニューを非表示(max/linuxのみ)
	Width       int    // 幅
	Height      int    // 高さ
	UserDataDir string // ユーザデータを再利用する(絶対パスで指定)
}

func (opt *DriverOption) FixDefault() {
}

// chrome driverを起動時、5秒待機その後、タイトルを返却
func sampleGetTitle(url string, opt *DriverOption) (string, error) {
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
	if err := runConsole(page, "hello world", &result); err != nil {
		log.Panic(err)
	}
	log.Println(result)

	time.Sleep(time.Second * 5)

	// タイトルの取得
	title, err := getTitle(page)
	if err != nil {
		return "", err
	}

	// select処理のサンプル
	sampleSelect(page)

	return title, nil
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
	if opt.UserDataDir != "" {
		args = append(args, fmt.Sprintf("--user-data-dir=%s", opt.UserDataDir))
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
// pageにスクリプトと変数を割り当てるだけで、動作はjavascriptまかせ
// よって、console.log('$word') にすると文字列扱いになり "$word" が出力される
// goでカバーするなら、argumentsは空でも良い
func runConsole(page *agouti.Page, word string, result *interface{}) error {
	body := "console.log($word);"
	arguments := map[string]interface{}{"$word": word}
	return page.RunScript(body, arguments, result)
}

func getTitle(page *agouti.Page) (string, error) {
	// agoutiの機能で、タイトルを取得 ----------
	title1, err := page.Title()
	if err != nil {
		return "", err
	}

	// goqueryを利用してタイトル取得 ----------
	// HTMlを取得
	html, err := page.HTML()
	if err != nil {
		return "", err
	}
	// htmlを解析してタイトル取得
	title2, err := analyzeHtml(html)
	if err != nil {
		return "", err
	}

	// それぞれの方法で違いがないかチェック ----------
	if title1 != title2 {
		return "", fmt.Errorf("not equal title. %v, %v", title1, title2)
	}

	return title1, nil
}

// qoqueryにより、htmlを解析してtitle取得
func analyzeHtml(html string) (string, error) {
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

// 要素取得サンプル
//
// ダイレクトにerrorチェックをする方法が無いため、countでエラーチェック(見つからないのも通常挙動だから？)
//
func sampleSelect(page *agouti.Page) {
	// find ----------
	// 一つだけ存在することを前提として検索(存在しない/複数存在するの両方でエラー)
	{
		// success(１つのみ存在するため成功)
		s := page.Find("title")
		if _, err := s.Count(); err != nil {
			log.Printf("error %v\n", err)
		}
	}

	{
		// error(複数存在するためエラー)
		s := page.Find("div")
		if _, err := s.Count(); err != nil {
			log.Printf("error %v\n", err)
		}
	}

	{
		// error(存在しないためエラー)
		s := page.Find("xyz")
		if _, err := s.Count(); err != nil {
			log.Printf("error %v\n", err)
		}
	}

	// First ----------
	{
		// success
		s := page.First("div")
		if _, err := s.Count(); err != nil {
			log.Printf("error %v\n", err)
		}
	}

	{
		// success
		s := page.First("input.gLFyf.gsfi")
		if _, err := s.Count(); err != nil {
			log.Printf("error %v\n", err)
		} else {
			// input value に値をセット
			s.Fill("mmmmmm")
		}
	}

	// All ----------
	{
		// success
		s := page.All("div")
		if _, err := s.Count(); err != nil {
			log.Printf("error %v\n", err)
		}
	}
}
