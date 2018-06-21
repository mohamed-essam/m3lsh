package m3lsh

import (
	"reflect"
)

// CatcherWrapper a wrapper for an exception catch function
type CatcherWrapper struct {
	typ reflect.Type
	fn  catchFn
}

// Catcher generate CatcherWrapper for a specific exception type with a handler
func Catcher(typ exception, fn catchFn) *CatcherWrapper {
	return &CatcherWrapper{typ: reflect.TypeOf(typ), fn: fn}
}
