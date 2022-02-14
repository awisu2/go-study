# switch-study

- switch は break を書かなくても次の case は実行しない
  - 複数条件が必要な場合は,で複数設定する

## sample

```go
package main

import "log"

func main () {
	y := sample(1)
	log.Println(y)  // 1

	y = sample(2)
	log.Println(y)  // 2

	y = sample(4)
	log.Println(y)  // 2

	y = sample(100)
	log.Println(y)  // 0
}

func sample(x int) int {
	y := 0
	switch x {
	case 2, 4:
		y = 2
	case 1:
		y = 1
	default:
		y = 0
	}
	return y
}

```
