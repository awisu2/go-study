package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func init() {
	// log setting
	log.SetPrefix("simple_sample: ")
	log.SetFlags(0)
}

func main() {
	fmt.Println("hello")
	log.Println("log hello")

	a, b := 9, 3
	c, err := division(a,  b)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d / %d = %d", a, b, c)
}

// error は最後の引数
func division(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("b can't 0")
	}

	// 四捨五入: math.Round(float64), 切り捨て: math.Floor(float64), 切り上げ: math.Ceil(float64)
	return int(math.Round(float64(a)  / float64(b))), nil
}
