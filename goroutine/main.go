package main

import (
	"fmt"
	"time"
)

func main() {
	sampleHello()
	sampleChannel()
	sampleChannelLimit()
}

func sampleHello() {
	// 場合によっては、worldは出力されない (goroutineの処理が終わる前にプログラムが終わる)
	go fmt.Println("world")
	fmt.Println("hello")
}

// sample Channel
// 値をチャンネルに送信できる

func sum(nums []int, ch chan int) {
	n := 0
	for _, num := range nums {
		n += num
	}
	// nをチャンネルに送信
	ch <- n
}

func sampleChannel() {
	// channel は　makeによって生成可能
	var ch chan int = make(chan int)
	go sum([]int{1, 2, 3}, ch)
	go sum([]int{4, 5, 6}, ch)
	go sum([]int{7}, ch)
	go sum([]int{8}, ch)
	go sum([]int{9}, ch)
	go sum([]int{10}, ch)
	// 結果は終わった順に格納される(ランダム)
	x, y, z, a, b, c := <-ch, <-ch, <-ch, <-ch, <-ch, <-ch

	fmt.Printf("%d, %d, %d, %d, %d, %d\n", x, y, z, a, b, c)
}

func dowble(n int, ch chan int) {
	ch <- n * n
	// close(ch)
}

func sampleChannelLimit() {
	// channelのバッファ可能数を2個に制限
	var ch chan int = make(chan int, 2)

	ch <- 1
	ch <- 2
	// 最大を超えるためここでエラー
	// fatal error: all goroutines are asleep - deadlock!
	// ch <- 3
	fmt.Printf("%d, %d\n", <-ch, <-ch)

	// 10回の入力を出力したら終了
	end := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		end <- 1
	}()

	i := 0
	for {
		// selectによってその処理が可能な場合にのみ処理される
		//
		// できる限りの速度でchに値を送信しているが、バッファ可能数は2のため、
		// 上記の出力処理がバッファ分を利用するまでdefaultになる
		// 例：アクセスしたいURLのリストを並べておき、それへのアクセスを実行(最大同時数の制限と、可能な限りの速度で実行することが可能)
		//
		select {
		case ch <- i:
			fmt.Printf("send %d\n", i)
			i++
		case <-end:
			return
		default:
		}
	}
}
