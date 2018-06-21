package m3lsh

import (
	"testing"
)

func TestFormatException(t *testing.T) {
	tex := &testException{}
	formatException(tex)
	if len(tex.stackTrace) == 0 {
		t.Error("No stack trace added")
	}
	emptyInfo := threadInfo{}
	if tex.thread == emptyInfo {
		t.Error("Thread info is null")
	}
}
