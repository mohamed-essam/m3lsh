package m3lsh

import (
	"reflect"
)

type tryFn func()
type catchFn func(interface{})

// Try catch all errors and exceptions and return them without failing
func Try(function tryFn) (err exception) {
	defer func() {
		if r := recover(); r != nil {
			err = recoveredException(r)
		}
	}()
	function()
	return nil
}

// TryCatch catch errors specified in catchers, if an *m3lsh.BaseException is given it takes any kind of exception, otherwise, exception Type must exactly match given type
func TryCatch(function tryFn, catchers ...*CatcherWrapper) (err exception) {
	defer func() {
		if r := recover(); r != nil {
			err = recoveredException(r)
			errType := reflect.TypeOf(err)
			// Find if any catchers fit
			for _, CatcherWrapper := range catchers {
				if isBaseException(CatcherWrapper.typ) || CatcherWrapper.typ == errType {
					CatcherWrapper.fn(err)
					return
				}
			}
			panic(err)
		}
	}()
	function()
	return nil
}

// Throw Format and throw error message with type and message
func Throw(ex exception, message string) {
	ex.setMessage(message)
	formatException(ex)
	panic(ex)
}
