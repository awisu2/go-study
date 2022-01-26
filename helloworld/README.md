# helloworld

install から helloworld の出力までを実行する

## 1: go のインストール

公式からインストールファイルをダウンロードして実行
[Download and install - The Go Programming Language](https://golang.org/doc/install)

## 2: helloworld モジュールの作成

```bash
mkdir myGoModule
cd myGoModule
go mod init {myGoModule}
```

## 3: main.go を作成

_main.go_

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world !")
}
```

## 4: 実行

`go run .`

## 5: もっと学ぶ

公式の非常に丁寧な練習環境があります
[A Tour of Go](https://go-tour-jp.appspot.com/welcome/1)
