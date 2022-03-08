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
					return ctx.Err()
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
		fmt.Printf("%v\n", err)
		return 0, err
	}

	return sum, nil
}
