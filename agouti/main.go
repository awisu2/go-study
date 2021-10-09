package main

import (
	"fmt"
	"os"

	"github.com/sclevine/agouti"
)

func main() {
	options:= agouti.ChromeOptions(
		"args", []string {
			"--headless",
			"--disable-gpu",
		},
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

    fmt.Println(page.Title())
}