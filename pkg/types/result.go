package types

type Result[T any] struct {
	Data  T
	Error error
}

type ResultOrError[T any] struct {
	Data  *T
	Error error
}
