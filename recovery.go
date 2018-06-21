package m3lsh

import (
	"reflect"
)

type tryFn func()
type catchFn func(interface{})

func Try(function tryFn) (err Exception) {
	defer func() {
		if r := recover(); r != nil {
			err = recoveredException(r)
		}
	}()
	function()
	return nil
}

func TryCatch(function tryFn, catchers ...*catcher) (err Exception) {
	defer func() {
		if r := recover(); r != nil {
			err = recoveredException(r)
			errType := reflect.TypeOf(err)
			// Find if any catchers fit
			for _, catcher := range catchers {
				if isBaseException(catcher.typ) || catcher.typ == errType {
					catcher.fn(err)
					return
				}
			}
			panic(err)
		}
	}()
	function()
	return nil
}

func Throw(ex Exception, message string) {
	ex.setMessage(message)
	formatException(ex)
	panic(ex)
}
