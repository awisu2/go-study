package main

import (
	"fmt"
	"runtime"
	"strconv"
)

const LINE = "----- "

func main() {
	fmt.Println("Hello, 世界")

	sampleLoop()
	sampleIf()
	sampleSwitch()
	sampleDefer()
	samplePointer()
	sampleStruct()
	sampleArray()
	sampleMaps()
}

func sampleLoop() {
	fmt.Println(LINE + "sampleLoop")

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// ; の省略も可能
	sum = 0
	for sum < 10 {
		sum += 1
	}
	fmt.Println(sum)

	// 無限ループ
	sum = 0
	for {
		sum += 1
		if sum >= 10 {
			break
		}
	}
	fmt.Println(sum)

	// slice は index または rangeでループが可能
	nums := []int{2, 4, 8}
	for i := 0; i < len(nums); i++ {
		fmt.Printf("index: %d, value: %d. ", i, nums[i])
	}
	fmt.Println()
	for i, v:= range nums {
		fmt.Printf("index: %d, value: %d. ", i, v)
	}
	fmt.Println()
	// 省略記述も可能(値がいらない場合)
	// for _, v:= range nums
	// for i, _:= range nums
	// for i := range nums // 後半の値が不要な場合は省略も可能
}

func sampleIf() {
	fmt.Println(LINE + "sampleIf")

	// 条件式内のスコープで宣言(b)
	a := "a"
	if b := a; b == "a" {
		fmt.Println("b is a!")
	}
	// bは使えない
	// fmt.Println(b)
}

func sampleSwitch() {
	fmt.Println(LINE + "sampleSwitch")

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func sampleDefer() {
	fmt.Println(LINE + "sampleDefer")

	// 関数の最後まで実行を遅延させる(ファイルのクローズ処理などに)
	defer fmt.Println("[defer] 1")

	// deferを複数宣言すると、最後のものから実行される(a, b, c と詰めると c, b, a の順になる)
	for i := 0; i < 3; i++ {
		defer fmt.Println("[defer] stack " + strconv.Itoa(i))
	}

	fmt.Println("[defer] 2")
}

func samplePointer() {
	fmt.Println(LINE + "samplePointer")

	i, j := 42, 10

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 5   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

type Vertex struct {
	X int
	Y int
}

func sampleStruct() {
	fmt.Println(LINE + "sampleStruct")

	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Println(v.X + v.Y)

	// pointer
	// structのポインタに値参照時 *p を p としてショートカットできる
	p := &v
	fmt.Printf("(*p).X = p.X <=> %d = %d\n", (*p).X, p.X)

	// フィールド指定での宣言も可能かつ、すべてを初期化する必要は無い
	v2 := Vertex{Y: 3}
	fmt.Println(v2)
}

func sampleArray() {
	fmt.Println(LINE + "sampleArray")

	// array (固定長)
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a, a[0], a[1])

	a2 := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("array(int)", a2)

	// slice (可変長, []T 配列数を空で宣言)
	var s []int = a2[1:4] // 1 <= i < 4 で取得(後者が未満であることに注意)
	fmt.Println("slice:", s)

	// slice は a2 の参照を持っており、元の配列の値が変化する
	s[0] = 99
	fmt.Println("change slice (array reference):", a2)

	// slice は array のリテラル
	//
	// これは配列
	// [3]bool{true, true, false}
	//
	// これはslice(配列を作成しそれを参照するsliceを返却)
	// []bool{true, true, false}

	// スライスでの取得方法(以下はすべて等価)
	l := len(a)
	fmt.Println(a[:], a[0:], a[:l], a[0:l])

	// len: スライスの長さ
	// cap: 配列の長さ(スライスの参照している)
	s2 := a2[0:1]
	_printLenCap(s2)

	s2 = s2[:] // 指定なしの上限だと len内で収まる
	_printLenCap(s2)

	s2 = s2[0:3] // cap内であれば、元のsliceの範囲を超えられる
	_printLenCap(s2)

	// append で要素を追加
	// sliceの範囲で追加されるため、元の要素の途中に挿入されることもある
	fmt.Println("before append", a2) // [2 99 5 7 11 13]
	s2 = append(s2, 88) 
	_printLenCap(s2)
	fmt.Println("after append", a2) // [2 99 5 88 11 13]

	// nil slice
	var nilSlice []int // 元となる値を持たない
	if nilSlice == nil { // nil 扱いになる
		fmt.Println("nilSlice is nil!")
	}

	// slice create by make()
	s3 := make([]int, 5) // len = 5, cap = 5
	_printLenCap(s3)

	s4 := make([]int, 3, 5) // len = 3, cap = 5
	_printLenCap(s4)

	// 多重配列も可能
	ss := [][]string {
		{"x", "x"},
		{"x", "x"},
	}
	ss[0][0], ss[1][1] = "o", "o"
	fmt.Println(ss)
}

func _printLenCap(s []int) {
	fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s)
}

func sampleMaps() {
	fmt.Println(LINE + "sampleMaps")

	m := make(map[string]Vertex)
	m["point 1"] = Vertex{99, 99}
	m["point 2"] = Vertex{55, 55}
	fmt.Println(m, m["point 1"]) 

	// 要素の取得と存在確認(okで確認)
	if v, ok := m["missing"]; !ok {
		// v には値無しで初期化された値がセットされる
		fmt.Println("no exists key", v)
	}

	// 存在しないキーを指定すると初期化した値が返却される(元のmapに変化はなし)
	fmt.Println("not change map. ", m)

	// literal
	m2 := map[string]int{
		"coffie": 300,
		"bread": 200,
	}
	fmt.Println(m2)

	// delete
	delete(m2, "coffie")
	delete(m2, "missing") // エラーにはならない
	fmt.Println(m2)
}