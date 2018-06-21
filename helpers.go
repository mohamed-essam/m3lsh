package m3lsh

import (
	"fmt"
	"reflect"
)

func implementsException(i reflect.Type) bool {
	return i.Implements(reflect.TypeOf((*Exception)(nil)).Elem())
}

func recoveredException(r interface{}) Exception {
	if !implementsException(reflect.TypeOf(r)) {
		// Warn as non m3lsh error
		fmt.Printf("Exception %s handled by m3lsh", r)
		err := &BaseException{Message: fmt.Sprintf("%v", r)}
		formatException(err)
		return err
	} else {
		return r.(Exception)
	}
}

func isBaseException(t reflect.Type) bool {
	return t == reflect.TypeOf(&BaseException{})
}
