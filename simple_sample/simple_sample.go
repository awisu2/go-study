package main

// 複数パッケージのインストール
import (
	"errors"
	"fmt"
	"log"
	"math"

	"go-study/simple_sample/sub"
)

// const は　character, string, bool, int　のみ
const MODULE_NAME = "simple_sample"

// 関数外でも var で宣言できる(":=" は不可)
var moduleName = MODULE_NAME

func init() {
	// log setting
	log.SetPrefix(moduleName + ": ")
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

	// call module
	fmt.Println(sub.Hello())

	// end
	log.Fatal("sample finish! (this is fake fatal)")
}

// 同じ型の場合、最後の宣言以外は省略できる
// public化: 先頭大文字
// error: 返却値の最後に設置
func Division(a , b int) (int, error) {
	if b == 0 {
		// create error
		return 0, errors.New("b can't 0")
	}

	// キャスト: T(v)
	// 四捨五入: math.Round(float64), 切り捨て: math.Floor(float64), 切り上げ: math.Ceil(float64)
	return int(math.Round(float64(a)  / float64(b))), nil
}
