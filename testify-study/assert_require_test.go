package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type ValueObject = struct {
	Value string
}

func TestAssert(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "want 123")
	assert.Equal(t, "hello", "hello", "want hello")

	// assert for nil (good for errors)
	obj := &ValueObject{Value: "Samething"}
	assert.Nil(t, obj)

	// assert for not nil (good when you expect something)
	if assert.NotNil(t, obj) {
		// now we know that object isn't nil
		assert.Equal(t, "Something", obj.Value)
	}
}

func TestRequire(t *testing.T) {
	// require
	require.Equal(t, 123, 234, "want 123")

	// not run because before test missing
	require.Equal(t, 123, 999, "want 123")
}
