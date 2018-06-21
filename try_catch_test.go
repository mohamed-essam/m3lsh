package m3lsh

import (
	"reflect"
	"testing"
)

func TestTryCatch(t *testing.T) {
	TryCatch(func() {
		Throw(&testException{}, "TEST")
	}, Catcher(&testException{}, func(ex interface{}) {
		if reflect.TypeOf(ex) != reflect.TypeOf(&testException{}) {
			t.Fatal("Wrong exception type")
		}
		e := ex.(*testException)
		if e.Message != "TEST" {
			t.Error("Exception has wrong message")
		}

		if len(e.Stacktrace()) <= 1 {
			t.Error("Stack trace too short")
		}
	}), Catcher(&BaseException{}, func(ex interface{}) {
		t.Error("Called wrong catcher")
	}))
}

func TestTryCatchPanic(t *testing.T) {
	TryCatch(func() {
		panic("NO")
	}, Catcher(&testException{}, func(ex interface{}) {
		t.Error("Called wrong catcher")
	}), Catcher(&BaseException{}, func(ex interface{}) {
		if reflect.TypeOf(ex) != reflect.TypeOf(&BaseException{}) {
			t.Fatal("Wrong exception type")
		}
		e := ex.(*BaseException)
		if e.Message != "NO" {
			t.Error("Exception has wrong message")
		}

		if len(e.Stacktrace()) <= 1 {
			t.Error("Stack trace too short")
		}
	}))
}

func TestTryCatchNoCatcher(t *testing.T) {
	defer func() {
		recover()
	}()
	TryCatch(func() {
		panic("WO?")
	}, Catcher(&testException{}, func(ex interface{}) {
		t.Fatal("Called wrong catcher")
	}))
	t.Fatal("Did not panic")
}
