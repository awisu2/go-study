package main

import (
	"fmt"
	"time"
)

func main() {
	sampleHello()
	sampleChannel()
	sampleChannelBuffer()
	// samplePingPong()
	// samplePingPong2()
}

func sampleHello() {
	go func() {
		fmt.Println("world")		
	}()
	fmt.Println("hello")
	time.Sleep(1* time.Second) // If don't sleep, world will not output

}

// SampleChannel()
// 
// Channels can exchanged inside and outside goroutine
//
// Waiting for the channel to respond will stop the process. 
// This means that if the channel does not return a response,
// it will not run any further and will not complete
func sampleChannel() {
	// Create channel
	var ch = make(chan int)

	sum := func(nums []int, ch chan int) {
		n := 0
		for _, num := range nums {
			n += num
		}
		ch <- n // send total to channel
	}
	go sum([]int{1, 2, 3}, ch)
	go sum([]int{4, 5, 6}, ch)
	go sum([]int{7, 8 , 9}, ch)

	// Stop until the channel return a response
	//
	// The order of the return values will be different each times,
	// because they are set in the order of processing completion.
	a := <-ch // wait 1th response
	b, c := <-ch, <-ch // wati 2,3th response

	fmt.Printf("%d, %d, %d\n", a, b, c)
}

func sampleChannelBuffer() {
	// channelのバッファ可能数を2個に制限
	var ch chan int = make(chan int, 2)

	ch <- 1
	ch <- 2
	// ch <- 3 // An error will occur, but this is not bad code because this is just blocking.

	fmt.Printf("%d, %d\n", <-ch, <-ch)

	// 上限付き同時処理実行
	work := func(f func(int) int, values []int, buffer int) chan int{
		running := make(chan bool, buffer)
		ch := make(chan int)
		go func() {
			n := 0

			for _, v := range values {
				// change scope
				v := v

				// job start. Block with buffer
				running <- true

				go func() {
					ch <- f(v)
					<- running // one job end

					// check comple
					if n++; n == len(values) {
						close(ch)
					}
				}()
			}
		}()

		return ch
	}

	calc := func(v int) int {
		// any times need this job
		time.Sleep(1 * time.Second)

		// complete
		return v * 2
	}

	ch = work(calc, []int{1,2,3,4,5}, 3)
	for i := range ch {
		fmt.Println(i)
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
	// channelをcloseすることで完了を通知
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