package m3lsh

import "fmt"

type BaseException struct {
	Message    string
	thread     threadInfo
	stackTrace []stackTraceLine
}

func (b BaseException) Stacktrace() []string {
	ret := make([]string, 1+len(b.stackTrace))
	ret[0] = b.thread.Format()

	for idx, x := range b.stackTrace {
		ret[idx+1] = x.Format()
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
	routineId    int64
	routineState string
}

func (t threadInfo) Format() string {
	return fmt.Sprintf("goroutine %d [%s]", t.routineId, t.routineState)
}

type stackTraceLine struct {
	file   string
	line   int64
	pkg    string
	method string
}

func (s stackTraceLine) Format() string {
	return fmt.Sprintf("%s in %s:%d in package %s", s.method, s.file, s.line, s.pkg)
}
