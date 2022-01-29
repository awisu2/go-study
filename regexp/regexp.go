package main

import (
	"log"
	"regexp"
)

// 拡張子にヒット
var ReExt = regexp.MustCompile(`.[^.]+$`)

func main() {
	match()
	find()

}

func match () {
	s := "image.jpg"

	// match!
	if ReExt.MatchString(s) {
		log.Println("match!")
	}
}

func find() {
	s := "image.jpg"
	_s := ReExt.FindString(s)
	log.Println(s, _s) // image.jpg .jpg
}

// 正規表現で文字列置換
func replace(src string, reg string, repl string) string{
	re := regexp.MustCompile(reg)
	return re.ReplaceAllString(src, repl)
}