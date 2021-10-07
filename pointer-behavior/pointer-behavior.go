package main

import (
	"fmt"
)

func main() {
	q1IsChangeArrayArg()
	q2IsChangeScaraArg()
	q3IsChangeRange()
	q4IsChangeStructArg()
	q5IsReturnStructIsPointer()
}

// array, slice はポインタにしなくても参照渡し
func q1IsChangeArrayArg() {
	arr := []int{0}
	q1FuncNoPointer(arr)
	fmt.Println(arr) // [99]. changed

	arr = []int{0}
	q1FuncPointer(&arr)
	fmt.Println(arr) // [99]. changed
}

func q1FuncNoPointer(arr []int) {
	arr[0] = 99
}

func q1FuncPointer(arr *[]int) {
	(*arr)[0] = 99
}

func q2IsChangeScaraArg() {
	i := 0
	q2FuncNoPointer(i)
	fmt.Println(i) // 0. no change

	i = 0
	q2FuncPointer(&i)
	fmt.Println(i) // 99. changed
}

func q2FuncNoPointer(i int) {
	i = 99
}

func q2FuncPointer(i *int) {
	*i = 99
}

func q3IsChangeRange() {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		v += 1 // this is new instance
	}
	fmt.Println(arr)  // [1 2 3]. no change

	arr = []int{1, 2, 3}
	for i := range arr {
		arr[i] += 1
	}
	fmt.Println(arr)  // [2 3 4]. changed
}

type User struct {
	Name string
}

// struct はpointerで渡さないと新規インスタンスとして渡される
func q4IsChangeStructArg() {
	user := User{Name: "jon"}
	q4FuncNoPointer(user)
	fmt.Println(user) // {jon}. no change

	user = User{Name: "jon"}
	q4FuncPointer(&user)
	fmt.Println(user) // {Smith}. changed
}

func q4FuncNoPointer(user User) {
	user.Name = "Smith"
}

func q4FuncPointer(user *User) {
	user.Name = "Smith"
}

var q5User User
func q5IsReturnStructIsPointer() {
	user := q5FuncReturnStruct()
	user.Name = "Smith"
	fmt.Println(q5User) // Jon. no change

	userP := q5FuncReturnStructPointer()
	userP.Name = "Smith"
	fmt.Println(q5User) // Smith. changed
}

func q5FuncReturnStruct() User {
	q5User = User{"Jon"}
	return q5User
}

func q5FuncReturnStructPointer() *User {
	q5User = User{"Jon"}
	return &q5User
}
