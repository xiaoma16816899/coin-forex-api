package utils

import (
	"errors"
	"reflect"
	"time"

	"server.com/pkg/utils/logger"
)

// Try calls the function and return false in case of error.
func Try[T any](callback func() (T, error)) (data T, panicErr any, err error) {
	defer func() {
		panicErr = recover()
	}()

	data, err = callback()
	return
}

func NewTryCatch[T any](main func() (T, error)) *TryCatch[T] {
	try := TryCatch[T]{mainFn: main}
	try.Exec()
	return &try
}

// Try Catch with chaining method, abstract types
type TryCatch[T any] struct {
	mainFn   func() (T, error)
	attempt  uint // default 0
	retries  uint // default 0
	panicErr any
	data     T
	err      error
}

// Exec CatchFn(), can do: check panic, err and retry one each implement
func (t *TryCatch[T]) Catch(catchFn func(try *TryCatch[T])) *TryCatch[T] {
	catchFn(t)
	return t
}

// Retry N times when mainFn return with err != nil,
func (t *TryCatch[T]) Retry(number uint) *TryCatch[T] {
	if t.err != nil {
		t.retries = number
		for t.attempt <= t.retries {
			t.Exec()
			if t.err == nil || !reflect.ValueOf(&t.data).Elem().IsZero() {
				break
			}
		}
	}

	return t
}

func (t *TryCatch[T]) Exec() {
	if t.attempt <= t.retries {
		t.attempt++
		t.data, t.panicErr, t.err = Try(t.mainFn)
	}
}

func (t TryCatch[T]) Value() (data T, panicErr any, err error) {
	return t.data, t.panicErr, err
}

// Just an example
func TryExample() any {
	number, panicErr, err := NewTryCatch(
		func() (int, error) {
			list := []int{1, 2, 3}
			return list[4], nil // will panic

		}).Catch( // catchError
		func(try *TryCatch[int]) {
			if try.panicErr != nil {
				try.data = 10000
			}
		}).Catch( // retry one more
		func(try *TryCatch[int]) {
			if try.err != nil {
				try.data, _ = try.mainFn()
			}
		}).Value() // .data or .err or .panicError

	i := 0
	NewTryCatch(func() (any, error) {
		i++
		logger.Debug("Atempts=>", i)
		time.Sleep(time.Second * 3)
		return nil, errors.New(("need trigger retry"))
	}).Retry(2)

	logger.Debug(number, panicErr, err)
	return number
}
