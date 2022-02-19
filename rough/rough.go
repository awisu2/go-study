package main

import "fmt"


const cns1 = 1
var str1 = "a"
var sct1 =  struct{
	Num int
} {
	Num: 1,
}
type ABC string
const (
	AbcA ABC ="a"
	AbcB ABC ="b"
	AbcC ABC ="c"
)

// interface, struct
type Duck = interface{
	GwaGwa()
}

type Human struct {
	name string
}

func (h *Human) GwaGwa() {
	fmt.Printf("hello I'm %s\n", h.name)
}

type Android struct {
	readyText string
	Human
}

func (r *Android) GwaGwa() {
	fmt.Printf("b-!! b-!! i'm %s, %s\n", r.name, r.readyText)
}

// sample vars
func vars() {
	fmt.Printf("%d\n", cns1)
	fmt.Printf("%s\n", str1)
	fmt.Printf("%v\n", sct1)

	fmt.Printf("%s\n", AbcA)

	str := "a"
	abc := AbcA
	fmt.Println(str, abc)
	
	// str = AbcA // cannot use AbcA (constant "a" of type ABC) as string value in assignment
	str = string(AbcA) // it's can. with convert.
	abc = "d" // it's can. without convert, not decraration value
	fmt.Println(str, abc)
}

// sample array
func printArray(arr [3]int) {
	fmt.Println(arr)
}

func printSlice(slice []int) {
	fmt.Println(slice)
}

func sampleArraySlice() {
	arr1 := [3]int{0, 1, 2}
	arr2 := [...]int{0, 1, 2}
	slice1 := []int{0, 1, 2, 3, 4, 5}
	slice2 := make([]int, 3) // [0 0 0]
	slice2[1] = 9

	// array
	printArray(arr1) // [0 1 2]
	printArray(arr2) // [0 1 2]
	// printArray([...]int{1, 2, 3, 4}) // error, not same length
	// printArray(slice1) // error

	// slice
	// printSlice(arr1) // error
	printSlice(slice1) // [0 1 2 3 4 5]
	printSlice(slice2) // [0 9 0]

	fmt.Println(len(arr1), len(slice1)) // 3 6

	// array to slice
	fmt.Println(arr1[:]) // [0 1 2]
	fmt.Println(arr1[1:]) // [1 2]

	// slice to slice
	fmt.Println(slice1[2:]) // [0 1 2 3 4 5]
	// fmt.Println(slice1[999:]) // panic: runtime error: slice bounds out of range [999:6]
	fmt.Println(slice1[:2]) // [0 1]
	fmt.Println(slice1[2:3]) // [2]
	fmt.Println(slice1[2:2]) // []

	// need convert to slice
	slice := append(arr1[:], 999)
	fmt.Println(slice)// [0 1 2 999]

	// range
	for i, v := range arr1 {
		fmt.Println(i, v)
	}
	for i, v := range slice1 {
		fmt.Println(i, v)
	}
}

func sampleMap() {
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}

	type S1 struct {
		Name string
	}
	s1 := S1{Name: "a"}
	s2 := S1{Name: "b"}
	map2 := map[S1]int{s1: 1, s2: 2}

	// get value
	if v, ok := map1["a"]; !ok {
		fmt.Println("not exists!")
	} else {
		fmt.Println(v) // 1
	}

	if v, ok := map2[s2]; !ok {
		fmt.Println("not exists!")
	} else {
		fmt.Println(v) // 2
	}

	// loop
	for k, v := range map2{
		fmt.Println(k.Name, v) // a 1, b 2
	}
}

// sample interface
func PlayOnStage(duck Duck) {
	duck.GwaGwa()
}

func InterfaceSample() {
	var jonDo = Human{
		name: "jondo",
	}
	var oldOne = Android{
		readyText: "count down 3... 2... 1...",
		Human: Human{
			name: "nanasi",
		},
	}

	PlayOnStage(&jonDo)
	PlayOnStage(&oldOne)
}

func sampleLoop() {
	// countup
	for i := 0; i < 3; i++ {
	}

	// while
	j := 0
	for j < 3 {
		j++
	}

	j = 0
	for {
		j++
		if j <= 3{
			break
		}
	}

	// goto
	for {
		goto exit
	}
	exit:
}

func main() {
	vars()
	InterfaceSample()
	sampleArraySlice()
	sampleMap()
	sampleLoop()
}