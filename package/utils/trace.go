package utils

import (
	"runtime"
	"time"
)

func GetFunctionName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	if frames == nil {
		return "error, can't get function name"
	}
	frame, _ := frames.Next()
	return frame.Func.Name()
}

func GetTimeUsage(startTime int64) int64 {
	return time.Now().UnixMicro() - startTime
}
