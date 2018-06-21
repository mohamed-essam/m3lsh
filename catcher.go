package m3lsh

import (
	"reflect"
)

type catcher struct {
	typ reflect.Type
	fn  catchFn
}

func Catcher(typ Exception, fn catchFn) *catcher {
	return &catcher{typ: reflect.TypeOf(typ), fn: fn}
}
