package main

import (
	"fmt"
	"log"
	"net/url"
)

type URL url.URL

var S = "https://foo.bar:1234/a/b/c?foo=bar&bar=foo#fragment"

func main() {
	parse()
	parse2()

	modify()
	copy() 
}

func parse() {
	_url, err := url.Parse(S)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(_url.String())		// https://foo.bar:1234/a/b/c?foo=bar&bar=foo#fragment
	fmt.Println(_url.Scheme)		// https
	fmt.Println(_url.Opaque)		// 
	fmt.Println(_url.User)			// 
	fmt.Println(_url.Host)			// foo.bar:1234
	fmt.Println(_url.Hostname())	// foo.bar
	fmt.Println(_url.Path)			// /a/b/c
	fmt.Println(_url.RawPath)		// 
	fmt.Println(_url.RawQuery)		// foo=bar&bar=foo
	fmt.Println(_url.Fragment)		// fragment
	fmt.Println(_url.Query())		// map[bar:[foo] foo:[bar]]
}

func parse2() {
	s := "scheme://userinfo@host/path?query#fragment"
	_url, err := url.Parse(s)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(_url.String())		// scheme://userinfo@host/path?query#fragment
	fmt.Println(_url.Scheme)		// scheme
	fmt.Println(_url.Opaque)		// 
	fmt.Println(_url.User)			// userinfo
	fmt.Println(_url.Host)			// host
	fmt.Println(_url.Hostname())	// host
	fmt.Println(_url.Path)			// /path
	fmt.Println(_url.RawPath)		// 
	fmt.Println(_url.RawQuery)		// query
	fmt.Println(_url.Fragment)		// fragment
	fmt.Println(_url.Query())		// map[query:[]]
}

func modify() {
	_url, err := url.Parse(S)
	if err != nil {
		log.Panic(err)
	}

	// プロパティを直接更新可能
	_url.Fragment = ""
	_url.RawQuery = "foo=foo&bar=bar"
	fmt.Println(_url)  // https://foo.bar:1234/a/b/c?foo=foo&bar=bar
}

// url.URLのコピー
//
// 一度、インスタンスの実態を引き渡しコピーを作成し、その後ポインタを取得いう手順
// copiedUrl := &(*_url) では元の参照に戻るため、実態コピーを作る１手が必要
//
func copy() {
	fmt.Println("copy -----")
	_url, _ := url.Parse(S)

	copy := func (_url url.URL) *url.URL {
		return &_url
	}
	copiedUrl := copy(*_url) // &(*_url) ではだめ
	copiedUrl.Host = "hoge.fuga"

	fmt.Println(_url) // https://foo.bar:1234/a/b/c?foo=bar&bar=foo#fragment
	fmt.Println(copiedUrl) // https://hoge.fuga/a/b/c?foo=bar&bar=foo#fragment
}