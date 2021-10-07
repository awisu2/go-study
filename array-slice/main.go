package main

import (
	"fmt"
	"sort"
)

func main() {
	sorting()
	sotingInts()
}


type User struct {
	Name string 
	Age int
}
func sorting() {
	// sort.Slice, sort.SliceStable
	// 
	// どちらも上書きソート。
	// SliceStableは安定ソートといい、同値の場合元のsliceの順序を保持する
	//
	slice := []User{
		{"A", 20},
		{"C", 10},
		{"B", 10},
		{"D", 15},
	}
	sort.Slice(slice, func(l, r int) bool {
		return slice[l].Age < slice[r].Age
	})
	fmt.Println(slice)

	slice = []User{
		{"A", 20},
		{"C", 10},
		{"B", 10},
		{"D", 15},
	}
	sort.SliceStable(slice, func(l, r int) bool {
		return slice[l].Age < slice[r].Age
	})
	fmt.Println(slice)
}

func sotingInts() {
	fmt.Println("sotingInts ----------")

	slice := []int{2,3,1}
	sort.Ints(slice)
	fmt.Println(slice)

	// reverse
	// 
	// sort.Interfaceを作成して、実行
	// sort.IntSliceは []int の再宣言
	// 
	reverse := sort.Reverse(sort.IntSlice(slice))
	sort.Sort(reverse)
	fmt.Println(slice)
}