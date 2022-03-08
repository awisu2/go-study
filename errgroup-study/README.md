# errgroup

- [errgroup package \- golang\.org/x/sync/errgroup \- pkg\.go\.dev](https://pkg.go.dev/golang.org/x/sync/errgroup)
- patterns
  - [Go Concurrency Patterns: Pipelines and cancellation \- The Go Programming Language](https://go.dev/blog/pipelines)

**NOTE**

- Can wait multi goroutine end. `g.Wait()`
- Can catch the error with a controlled goroutine error. `err := g.Wait()`
  - If want to close the goroutine on error, need `<- ctx.Done()`
    - at the time `ctx.Err()` is *context canceled*

## samples

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func sumWithMulti(ctx context.Context, num int, badnumber int, buffer int) (int, error) {
	g, ctx := errgroup.WithContext(ctx)

	// set values
	// バッファが足りない時セット処理で停止してしまうためGoroutine内でセット
	pip := make(chan int, num)
	g.Go(func() error {
		for i := 0; i < num; i++ {
			n := i + 1

			if n == badnumber {
				return fmt.Errorf("%d is bad number", n)
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case pip <- n:
				fmt.Printf("set %d\n", n)
			}
		}
		close(pip)
		return nil
	})

	// calc
	sum := 0
	mu := sync.Mutex{}
	for i := 0; i < buffer; i++ {
		g.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return nil
				case n, ok := <-pip:
					if !ok {
						return nil
					}
					fmt.Printf("get %d\n", n)
					mu.Lock()
					sum += n
					mu.Unlock()

					// any time cost
					time.Sleep(1 * time.Second)
				}
			}
		})
	}

	// wait: We can get error multi Go
	if err := g.Wait(); err != nil {
		return 0, err
	}

	return sum, nil
}
```
