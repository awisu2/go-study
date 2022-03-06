package main

import (
	"github.com/stretchr/testify/mock"
)

// our faunction
type Something = interface {
	Hello(name string) (string, error)
}

func Hello(something Something, name string) (string, error) {
	return something.Hello(name)
}

// mock (tipycaly write under mock package)
type SomethingMock struct {
	mock.Mock
}

func (m *SomethingMock) Hello(name string) (string, error) {
	args := m.Called(name)
	return args.String(0), args.Error(1)

}
