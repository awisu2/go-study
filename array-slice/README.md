# array-slice

## sort

SliceStable: 安定ソート の動作条件が不明。なので実装時利用するのであれば気をつける

- sort package を利用する
  - `func sort.Slice(x interface{}, less func(i int, j int) bool)`
  - 安定ソート: `func sort.SliceStable(x interface{}, less func(i int, j int) bool)`
    - [sort - Go 言語](https://xn--go-hh0g6u.com/pkg/sort/#SliceStable)
    - なんか元の slice の配列順を残してくれる。例だと年ソートの前の名前ソート順が反映される
- 基本上書きソート(php や python で言う破壊的ソート)
- 他にも特定の型に合わせた関数が用意されている
  - int: `sort.Ints(slice)`
  - 逆ソート: `sort.Sort(sort.Reverse(sort.IntSlice(slice)))`
    - sort.IntSlice でキャストして、interface を合致させている
- 逆ソートの例を見るとわかるが、もともと sort.Interface に適合させて実行という形だったらしい
  - sort.Sort が出てきたら確認しておくといいかも

## sample

```go
func sorting() {
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
}
```
