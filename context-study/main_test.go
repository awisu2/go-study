package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextValue(t *testing.T) {
	ctx := context.Background()

	key := "number"
	ctx = context.WithValue(ctx, key, 123)

	assert.Equal(t, ctx.Value(key), 123, "want 123")
	assert.Nil(t, ctx.Value("text"), "want 123")
}

func TestAnyMain(t *testing.T) {
	anyMain()
}

func TestContextDone(t *testing.T) {
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

func TestContextTimeout(t *testing.T) {
	contextTimeout()
}
