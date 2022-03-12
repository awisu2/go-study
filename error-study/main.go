package main

import (
	"errors"
	"fmt"
)

type ErrorA string

func (err *ErrorA) Error() string{
	return string(*err)
}

func (err *ErrorA) Is(target error) bool {
	return err.Error() == target.Error()
}

// It can run with errors
type ErrorWithUnwrap struct {
	detail string
	child error // wrap error
}

func (err *ErrorWithUnwrap) Error() string{
	return err.detail
}

func (err *ErrorWithUnwrap) Is(target error) bool {
	_target, ok := target.(*ErrorWithUnwrap)
	if !ok {
		return false
	}

	return err.Error() == target.Error() && err.child == _target.child
}

func (err *ErrorWithUnwrap) Unwrap() error{
	return err.child
}

func NewErrorA(detail string) error{
	err := ErrorA(detail)
	return &err
}

func NewErrorWithUnwrap(detail string, child error) error{
	return &ErrorWithUnwrap{detail, child}
}

func asIsError01() {
	var errA = NewErrorA("errA")

	var _errA *ErrorA
	if errors.As(errA, &_errA) {
		fmt.Println("ok. hitted. " + _errA.Error())
	} else {
		fmt.Println("ng")
	}
}

func asIsError02() {
	// wrap errA
	var errA = NewErrorA("errA")
	var errorWithUnwrap = NewErrorWithUnwrap("errWithUnwrap", errA)

	// as
	var _errWithUnwrap *ErrorWithUnwrap
	if errors.As(errorWithUnwrap, &_errWithUnwrap) {
		fmt.Println("ok. " + _errWithUnwrap.Error())
	} else {
		fmt.Println("ng")
	}

	var _errA *ErrorA
	if errors.As(errorWithUnwrap, &_errA) {
		fmt.Println("ok. " + _errA.Error())
	} else {
		fmt.Println("ng.")
	}

	// Is
	var _errWithUnwrap2 = NewErrorWithUnwrap("errWithUnwrap", errA)
	if errors.Is(errorWithUnwrap, _errWithUnwrap2) {
		fmt.Println("ok. " + _errWithUnwrap2.Error())
	} else {
		fmt.Println("ng")
	}

	var _errA2 = ErrorA("errA")
	if errors.Is(errorWithUnwrap, &_errA2) {
		fmt.Println("ok. " + _errA2.Error())
	} else {
		fmt.Println("ng.")
	}
}

func main() {
	asIsError01()
	asIsError02()
}