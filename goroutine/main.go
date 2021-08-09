package main

import (
	"fmt"
	"time"
)

func main() {
	// sampleHello()
	// sampleChannel()
	// sampleChannelLimit()
	// samplePingPong()
	samplePingPong2()
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


type PingPong struct {
	Hit chan int
	End chan bool
	Name string
	IsEnd bool
}

func (pingPong *PingPong) Lary(partner *PingPong) {
	defer pingPong.Close()

	for {
		// 基本的には責任範囲を自身のinstanceにとどめておく
		select {
		case v := <- pingPong.Hit:
			if !pingPong.IsEnd {
				time.Sleep(time.Second)
				v += 1
				fmt.Println(pingPong.Name, ":", v)
				partner.Hit <- v
			}
		case <- pingPong.End:
			fmt.Println(pingPong.Name, ": end")
			// 終了処理
			//
			// 以下の手順
			// 1. まずは自分が処理を停止し、相手に通知
			// 2. 相手はそれを受けて処理を停止し、それを再通知、自身は終了(return)
			// 3. 自身はすでに停止しているので、終了(return)
			//
			// 課題：
			// channelの停止の前にそれぞれ相手へ送らないことを確定しないと, send closed channelのエラーになる
			// 思った以上にchannnelのクローズはシビア。フラグをsync.Mutexする程度ではタイミングが合わない
			// またchannelを閉じないままgoroutineを閉めると、all goroutine asleep エラーになる
			//
			// 「selectにより、caseのどれか一つしか動作しないことを利用して終了状態を確定」している(IsEnd)
			//　私停止しましたを安全に設定できる状態で相互に問題のない状態にしそれぞれ終了を行う
			//
			// 今回はpingPongでコンパクトに纏めたかったので相互的な確認にしたが、
			// 上位の管理structなどを作って管理したほうが、有用なパターンは多いと思われる
			// ただし、全てを管理struct経由になるため各処理をgoroutineに分けた意味がなくならないように
			if pingPong.IsEnd {
				return
			} else {
				pingPong.IsEnd = true
				partner.End <- true
				if partner.IsEnd {
					return
				}
			}
		}
	}
}

func (pingPong *PingPong) Close() {
	fmt.Println(pingPong.Name, ": close")
	close(pingPong.Hit)
	close(pingPong.End)
}

func CreatePingPong(name string) *PingPong {
	// なるべく上限は指定したほうが良さそう
	// TOOD: ちょっと不明な点が多い
	// INFO: 両方のchannelを制限掛けずにセットしたらEndへ送信した時点で停止する問題があった
	// INFO: endのlimitを設定せず、pingPongの回数を多くしたらcloseしなかった(最後の待機時間を多くしてもだめ)
	// INFO: endのlimitを1に設定したらpingPongの回数を多くしてもcloseした
	// INFO: hitの方は設定してもしなくても、動作時間などに関係なく、普通に動作した。
	// INFO: endの型をint/boolと切り替えたがそれによっての動作は変わらなかった
	hit := make(chan int)
	end := make(chan bool, 1)
	return &PingPong {
		Hit: hit,
		End: end,
		Name: name,
	}
}

func samplePingPong() {
	// channel は　makeによって生成可能
	{
		ping, pong := CreatePingPong("ping"), CreatePingPong("pong")
		go ping.Lary(pong)
		go pong.Lary(ping)

		ping.Hit <- 0

		time.Sleep(time.Second * 5)
		ping.End <- true

		// 待機しないとerrorログを華麗にスルーしてしまうので待機
		time.Sleep(time.Second * 3)
	}
}

// 集中管理型
type PingPongController struct {
	Resist chan *PingPongPlayer
	End chan bool
	Hit chan *PingPongHit

	Players map[*PingPongPlayer]bool
}

func (c *PingPongController) Run() {
	defer c.Close()

	L: for {
		select {
		case player := <- c.Resist:
			fmt.Printf("resist %s\n", player.Name)
			c.Players[player] = true

		case hit := <- c.Hit:

			fmt.Printf("Hit %s %d\n", hit.Player.Name, hit.Num)
			for player := range c.Players {
				if (player != hit.Player) {
					player.Hitted <- hit.Num
					break
				}
			}

		case <- c.End:
			for player := range c.Players {
				c.UnResist(player)
			}
			break L
		}
	}
}

func (c *PingPongController) Close() {
	fmt.Println("PingPongController Close")
	close(c.Resist)
	close(c.End)
}

func (c *PingPongController) UnResist(player *PingPongPlayer) {
	fmt.Printf("UnResist %s\n", player.Name)
	delete(c.Players, player)
	// player.End <- true
	player.Close()
}

type PingPongPlayer struct {
	Controller *PingPongController
	Name string
	Hitted chan int
	End chan bool
	isEnd bool
}

func (p *PingPongPlayer) Run() {
	defer fmt.Println("PingPongPlayer Run End. " + p.Name)
	
	L: for {
		select {
		case n, ok := <- p.Hitted:
			if !ok {
				fmt.Println("HItted close " + p.Name)
				break L
			}
			fmt.Println("HItted " + p.Name)
			if n > 0 {
				time.Sleep(time.Second)
			}
			// close ではタイミングが合わない事があるため、フラグ管理も同時に行う
			// channelを実行中にcloseをするとループが回ってこないみたい
			// TODO: 無駄なフラグな気もするのでもっといい方法があればそっちを使いたい
			// (PingPongPlayer Run End が出力されない)
			if p.isEnd {
				fmt.Println("isEnd " + p.Name)
				break L
			}
			p.Controller.Hit <- &PingPongHit{
				Player: p,
				Num: n + 1,
			}

		case <- p.End:
			fmt.Printf("end %s\n", p.Name)
			break L
		}
	}
}

func (p *PingPongPlayer) Close() {
	fmt.Println("PingPongPlayer Close. " + p.Name)
	close(p.Hitted)
	close(p.End)
	p.isEnd = true
}

type PingPongHit struct {
	Player *PingPongPlayer
	Num int
}

func CreatePingPongController() *PingPongController{
	return &PingPongController{
		Resist: make(chan *PingPongPlayer),
		End: make(chan bool),
		Hit: make(chan *PingPongHit),
		Players: map[*PingPongPlayer]bool{},
	}
}

func CreatePingPongPlayer(controller *PingPongController, name string) *PingPongPlayer{
	return &PingPongPlayer{
		Controller: controller,
		Name: name,
		Hitted: make(chan int),
		End: make(chan bool),
	}
}

func samplePingPong2() {
	controller := CreatePingPongController()

	ping := CreatePingPongPlayer(controller, "ping")
	pong := CreatePingPongPlayer(controller, "pong")
	pang := CreatePingPongPlayer(controller, "pang")

	go controller.Run()
	go ping.Run()
	go pong.Run()
	go pang.Run()

	controller.Resist <- ping
	controller.Resist <- pong
	controller.Resist <- pang

	// ラリー
	ping.Hitted <- 0
	time.Sleep(time.Second * 5)
	controller.End <- true

	// 終了処理待ち
	time.Sleep(time.Second)
}