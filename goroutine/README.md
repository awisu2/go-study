# goroutine

[goroutine](https://go-tour-jp.appspot.com/concurrency/1)

- goroutine の挙動
  - いつ終わるかわからない
  - 無限ループにするなら停止させる機構が必要(context など)
- チャンネル
  - goroutine の外から中、中から外のどちらにも信号を送れる
  - バッファ
    - 空のときチャンネルの受信をブロック
    - 数が超えるときブロック
  - An error will occur if channel is blocked and goroutine does not exist
    - stop main process + no goroutine = process can't run more.
- select の活用
  - バッファ持ちのチャンネルへの入力と default で、上限だった場合は弾く
  - 複数のチャンネルを待機する daemon 処理

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

### 上限付き同時実行

- 処理を走らせる直前に、buffer 付きのチャンネルに値をセットし処理数をブロック
  - 個々の処理が終わったところで 1 つ減らす
  - このチャンネルは処理数管理のためなので、あまり意味ある値をもたせない
    - カウンターをもたせての終了処理もだめ、終了順が異なるから
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
