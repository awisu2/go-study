package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanStop(t *testing.T) {
	ctx := context.Background()
	err := cleanStop(ctx)

	assert.Nil(t, err, "clean Stop")

}
