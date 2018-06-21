package m3lsh

import (
	"testing"
)

type testException struct {
	BaseException
}

func TestTry(t *testing.T) {
	ex := Try(func() {
		Throw(&testException{}, "TEST")
	})

	if ex == nil {
		t.Log("Exception not returned")
		t.FailNow()
	}

	if ex.(*testException).Message != "TEST" {
		t.Log("Exception has wrong message")
		t.Fail()
	}

	if len(ex.Stacktrace()) <= 1 {
		t.Log("Stack trace too short")
		t.Fail()
	}
}
