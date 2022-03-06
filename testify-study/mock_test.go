package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMock(t *testing.T) {
	somethingMock := new(SomethingMock)

	// set function behavior
	name := "World"
	somethingMock.On("Hello", name).Return(fmt.Sprintf("Hello %v", name), nil)

	// run with mock
	//
	// note: The argument must be equivalent to the set value
	res, err := Hello(somethingMock, name)

	// test
	require.Nil(t, err, "not error happen")
	assert.Equal(t, res, "Hello World", "want Hello world")
}
