package m3lsh

import (
	"reflect"
	"testing"
)

type testNonException struct {
	randomField string
}

func TestImplementsException(t *testing.T) {
	if implementsException(reflect.TypeOf(&testNonException{})) {
		t.Fatal("testNonException does not implement exception")
	}
}

func TestRecoveredException(t *testing.T) {
	// when exception is actual thrown exception
	ex := &testException{}
	ex.setMessage("TEST MESSAGE")
	formatException(ex)
	recovered := recoveredException(ex)
	if recovered != ex {
		t.Error("Recovered exception does not match exception")
	}

	// when exception is string
	ex2 := recoveredException("TEST MESSAGE 2")
	if reflect.TypeOf(ex2) != reflect.TypeOf(&BaseException{}) {
		t.Fatal("Returned value is not base exception")
	}
	ex2C := ex2.(*BaseException)
	if ex2C.Message != "TEST MESSAGE 2" {
		t.Error("Returned exception message is wrong")
	}
}

func TestIsBaseException(t *testing.T) {
	tex := &testException{}

	if isBaseException(reflect.TypeOf(tex)) {
		t.Error("test exception is not base exception")
	}

	bex := &BaseException{}

	if !isBaseException(reflect.TypeOf(bex)) {
		t.Error("BaseException is BaseException (duh)")
	}
}
