package m3lsh

import (
	"runtime/debug"
	"strconv"
	"strings"
)

type Exception interface {
	Stacktrace() []string
	setThread(threadInfo)
	setStack([]stackTraceLine)
	setMessage(string)
}

func formatException(ex Exception) {
	vals := strings.Split(string(debug.Stack()), "\n")

	routineInfo := strings.Split(vals[0], " ")
	routineId, _ := strconv.ParseInt(routineInfo[1], 10, 64)
	routineState := routineInfo[2][1 : len(routineInfo[2])-2]

	lines := make([]stackTraceLine, len(vals)/2)

	for idx, line := range vals {
		if idx == 0 || idx+1 == len(vals) {
			continue
		}
		if idx%2 == 1 { // package and method
			lines[idx/2] = stackTraceLine{}
			if strings.Contains(line, ".") {
				lines[idx/2].pkg = strings.Split(line, ".")[0]
				lines[idx/2].method = strings.Split(strings.Split(line, ".")[1], "(")[0]
			} else {
				lines[idx/2].method = strings.Split(line, "(")[0]
				lines[idx/2].pkg = "golang"
			}
		} else {
			splitLine := strings.Split(strings.Trim(line, " \t"), " ")[0]
			lines[idx/2-1].file = strings.Split(splitLine, ":")[0]
			lines[idx/2-1].line, _ = strconv.ParseInt(strings.Split(splitLine, ":")[1], 10, 64)
		}
	}

	thread := threadInfo{routineId: routineId, routineState: routineState}

	ex.setStack(lines)
	ex.setThread(thread)
}
