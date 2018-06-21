package m3lsh

import (
	"fmt"
	"reflect"
)

func implementsException(i reflect.Type) bool {
	return i.Implements(reflect.TypeOf((*exception)(nil)).Elem())
}

func recoveredException(r interface{}) exception {
	if !implementsException(reflect.TypeOf(r)) {
		// Warn as non m3lsh error
		fmt.Printf("exception %s handled by m3lsh\n", r)
		err := &BaseException{Message: fmt.Sprintf("%v", r)}
		formatException(err)
		return err
	}
	return r.(exception)
}

func isBaseException(t reflect.Type) bool {
	return t == reflect.TypeOf(&BaseException{})
}
