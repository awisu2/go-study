package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"
)


func main() {
	// sampleFlag()
	// sampleFlagWithSetting()

	args, err := sampleArgParse()
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(args)
}

func sampleFlag() {
	// 後述のparse設定がない状態で, -n などの引数を受け取ると宣言無しエラーになる
	// flag provided but not defined: -n
	flag.Parse()
	log.Println(flag.Args())
	// 範囲外の場合は空文字返却
	log.Println(flag.Arg(0), flag.Arg(1), flag.Arg(2))
}

func sampleFlagWithSetting() {
	log.Println("---------- sampleFlagWithSetting")

	// parse設定がされている状態で、-xなど未指定の値を受け取るとusageが表示される --helpでも表示される
	//
	// Usage of /tmp/go-build2300064911/b001/exe/argparse:
	// 	-b    this is bool
	// 	-n int
	// 				this is num (default 999)
	// 	-s string
	// 				this is string (default "---")
	// 	-v    this is bool is set
	// exit status 2
	//
	// それぞれの型で値が取得できる
	var (
		// 引数は name, value, usage で、nameに - をつけたものが引数に受け渡される
		// 例: `go run . -n 123`
		i = flag.Int("n", 999, "this is num")
		s = flag.String("s", "---", "this is string")
		b = flag.Bool("b", false, "this is bool")
	)
	var p bool
	flag.BoolVar(&p, "v", false, "this is bool is set")
	flag.Parse()

	log.Println(*i, *s, *b, p)
}

// structを用意してそれを返却させる(Parse時の値を別途宣言する必要上、parserだけ関数内に分離して簡単に引数取得というのが難しいため)
type Args struct {
	S string
	C string
}

// map や []string でまとめないのは、宣言された引数としてアクセスできないから
// type Color string もやってみたが、argparseで利用するため[]string に変換する必要があり、コストが掛かる
const (
	RED = "red"
	BLUE = "blue"
	YELLOW = "YELLOW"
)
var Colors = []string{
	RED,
	BLUE,
	YELLOW,
}

func ValidateStringNotEmpty(s string, text string) error {
	if s == "" {
		return errors.New(text)
	}
	return nil
}

func ValidateStringsNotEmpty(strs []string, text string) error {
	for _, str := range strs {
		err := ValidateStringNotEmpty(str, text)
		if err != nil {
			return err
		}
	}
	return nil
}

func sampleArgParse() (*Args, error) {
	parser := argparse.NewParser("command", "parse description")

	// 見た目のわかりやすさのために関数を分離しているが、通常は同一関数内のほうがわかりやすい気がする
	s, c := argparserSetting(parser)

	if err := parser.Parse(os.Args); err != nil {
		return nil, err
	}

	return &Args{
		S: *s,
		C: *c,
	}, nil
}

// 取得する各引数の設定
func argparserSetting(parser *argparse.Parser) (*string, *string){
	var s = parser.String("s", "str", &argparse.Options{
		Required: true,
		// 空文字の場合はエラー
		Validate: func (strs []string) error{
			if err := ValidateStringsNotEmpty(strs, "[-s, --s] can't empty"); err != nil {
				return err
			}
			return nil
		},
		Help: "this is string sample",
	})

	var c = parser.Selector("c", "color", Colors, &argparse.Options{
		Default: string(RED),
		Help: "this is selector sample",
	})

	return s, c
}