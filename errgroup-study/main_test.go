package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiRun(t *testing.T) {
	sum, err := sumWithMulti(context.Background(), 10, 99, 3)

	if assert.Nil(t, err, "no error") {
		assert.Equal(t, sum, 55, "want 55")
	}
}

func TestMultiRunError(t *testing.T) {
	_, err := sumWithMulti(context.Background(), 10, 5, 3)

	assert.NotNil(t, err, "no error")
}
