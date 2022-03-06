package main

import (
	"context"
	"fmt"
	"time"
)

func contextTimeout() {
	var ctx, ctxTimeout context.Context

	ctx = context.Background()
	// ctxTodo := context.TODO()

	ctxTimeout, cancel := context.WithTimeout(ctx, 1 * time.Microsecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second): // not run
		fmt.Println("overslept")
	case <-ctx.Done(): // not run
		fmt.Println("ctx", ctx.Err()) 
	case <-ctxTimeout.Done(): // run this
		fmt.Println("ctxTimeout", ctx.Err())
	}
}

func contextDone(runCancel2 bool) {
	ctx := context.Background()
	ctxCancel1, cancel1 := context.WithCancel(ctx)
	ctxCancel2, cancel2 := context.WithCancel(ctxCancel1)
	defer func() {
		cancel1()
		cancel2()
	}()

	runGoroutine := func() <- chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select{
				case <- ctx.Done():
					fmt.Println("done base") // not run this
				case <- ctxCancel1.Done():
					fmt.Println("done cancel1")
				case <- ctxCancel2.Done():
					fmt.Println("done cancel2")
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	for range runGoroutine() {
		if runCancel2 {
			cancel2()
		} else {
			cancel1()
		}
		break
	}

	// for check ctx.Done()'s print
	time.Sleep(1 * time.Second)
}

func main() {
	contextTimeout()

	fmt.Println("run cancel 1 ---------")
	contextDone(false)
	// done cancel1
	// done cancel1
	// ...
	// done cancel2

	fmt.Println("run cancel 2 ---------")
	contextDone(true)
	// done cancel2

}