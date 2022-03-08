# goroutine

- [goroutine](https://go-tour-jp.appspot.com/concurrency/1)
- Can be safely closed when combined with conetxt: [context](../context-study/README.md)
- Can be convinient use multi goroutine with errgroup: [errgroup](../errgroup-study/README.md)

**NOTE**

- goroutine の挙動
  - いつ終わるかわからない
  - 無限ループにするなら停止させる機構が必要(context など)
- チャンネル
  - goroutine の外から中、中から外のどちらにも信号を送れる
  - buffer
    - 空のときチャンネルの受信をブロック
    - 数が超えるときブロック
  - non buffer channel. I think it same behavior buffer 1 channel. (`make(chan {any})` == `make(chan {any}, 1)`)
  - An error will occur if channel is blocked and goroutine does not exist
    - stop main process + no goroutine = process can't run more.
- select の活用
  - バッファ持ちのチャンネルへの入力と default で、上限だった場合は弾く
  - 複数のチャンネルを待機する daemon 処理
- 各種パターンでの利用
  - 上限付き同時処理(処理数固定式): 処理直前に動作数管理チャンネルに値をセット完了時に受信。これに対しforを行う。サンプルは下記。
    - errgroupを利用したpiplineパターンでの実装もある。完了条件の判断がチャンネルでできるのでこちらのほうがわかりやすいか
    - リストへ add することで適時処理を実行する daemon
  - cleanStop: daemon化したgoroutineを正常に終了させるサンプル(サンプルを後述)
    - 単に適当なチャンネルを作ってcloseするだけ
    - contextのcancelはいわゆる例外処理のため, 正常なクローズは自前で用意する必要あり

## samples

```go
// run goroutine
go func(){}()

// channel
ch := make(chan int)
ch := make(chan int, 3) // with buffer
ch <- 5 // send to channel
<- ch // wait channel response
v := <- ch // wait channel response and get value

// with range
for range ch {}
for i := range ch {}
```

## cleanStop

実際には呼び出し側に停止用チャンネルを返却するなどになると思われる

```go
// ただcloseするだけのチャンネルを作って完了時にcloseするだけ
// contextのCancelはいわゆる例外処理的なものなので、正常なgoroutineの完了処理は自分で作る必要がある
func cleanStop(ctx context.Context) error {
	// channel for stop
	stop := make(chan bool)
	g, ctx := errgroup.WithContext(ctx)

	// 複数のgoroutineが動作しても大丈夫
	for i := 0; i < 10; i++ {
		g.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-stop:
					return nil
				}
			}
		})
	}

	go func() {
		time.Sleep(1 * time.Second)
		// stop
		close(stop)
	}()

	return g.Wait()
}
```

### 上限付き同時実行(処理数固定式)

- 処理を走らせる直前に、buffer 付きのチャンネルに値をセットし処理数をブロック
  - 個々の処理が終わったところで 1 つ減らす
  - このチャンネルは処理数管理のためなので、あまり意味ある値をもたせない
    - カウンターをもたせての終了管理は可能(取り出し順により順に値が取得できるため)
- 完全終了チェックを個々の処理が終わったところで行う
  - deadlock 回避
  - ただこの処理は冗長なので、どこで処理するか？という問題はある
    - 外部の呼び出し側で処理をすることも可能

```go
// 上限付き同時処理実行
work := func(f func(int) int, values []int, buffer int) chan int{
  running := make(chan bool, buffer)
  ch := make(chan int)
  var mu sync.Mutex
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
        mu.Lock()
        n++
        mu.Unlock()
        if n == len(values) {
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
```
