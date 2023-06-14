package util

import (
	"fmt"
	"runtime"
)

func GetGoroutineID() int64 {
	buf := make([]byte, 64)
	n := runtime.Stack(buf, false)
	id := int64(0)
	fmt.Sscanf(string(buf[:n]), "goroutine %d ", &id)
	return id
}
