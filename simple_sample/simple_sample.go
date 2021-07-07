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
	// `:=`: 宣言と値のセットを同時に行う(`var msg string` & `msg = "hello"`)
	msg := "hello"

	fmt.Println(msg)
	log.Printf("log %s\n", msg)

	// slice: `var mySlice []int` array: `var myArray[3] int`
	// sliceのほうが扱いやすいとのこと
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	nums = append(nums, 0)

	for i := 0; i < len(nums) - 1; i++ {
		c, err := Division(nums[i],  nums[i+1])
		if err != nil {
			// go's log level just two. (print or Fatal)
			// log.Fatal(err)
			log.Println(err)
		} else {
			log.Printf("%d / %d = %d", nums[i], nums[i+1], c)
		}
	}


	// end
	log.Fatal("sample finish! (this is fake fatal)")
}

// public: 先頭大文字s
// error: 最後の引数
func Division(a int, b int) (int, error) {
	if b == 0 {
		// create error
		return 0, errors.New("b can't 0")
	}

	// 四捨五入: math.Round(float64), 切り捨て: math.Floor(float64), 切り上げ: math.Ceil(float64)
	return int(math.Round(float64(a)  / float64(b))), nil
}
