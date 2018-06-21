package m3lsh

import "fmt"

// BaseException parent of all exceptions, implements some basic methods and acts as a catch-all exception type
type BaseException struct {
	Message    string
	thread     threadInfo
	stackTrace []stackTraceLine
}

// Stacktrace get printable stack trace lines
func (b BaseException) Stacktrace() []string {
	ret := make([]string, 1+len(b.stackTrace))
	ret[0] = b.thread.format()

	for idx, x := range b.stackTrace {
		ret[idx+1] = x.format()
	}

	return ret
}

func (b *BaseException) setThread(thread threadInfo) {
	b.thread = thread
}

func (b *BaseException) setStack(stack []stackTraceLine) {
	b.stackTrace = stack
}

func (b *BaseException) setMessage(message string) {
	b.Message = message
}

type threadInfo struct {
	routineID    int64
	routineState string
}

func (t threadInfo) format() string {
	return fmt.Sprintf("goroutine %d [%s]", t.routineID, t.routineState)
}

type stackTraceLine struct {
	file   string
	line   int64
	pkg    string
	method string
}

func (s stackTraceLine) format() string {
	return fmt.Sprintf("%s in %s:%d in package %s", s.method, s.file, s.line, s.pkg)
}
