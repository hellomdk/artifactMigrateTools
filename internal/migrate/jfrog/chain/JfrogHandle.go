package chain

import (
	"strings"
)

type JfrogHandle struct {
	handle Handle
}

func (jh *JfrogHandle) Handle(content string) bool {
	if strings.Contains(content, ".jfrog") {
		return true
	}
	return jh.next(jh.handle, content)
}

func (jh *JfrogHandle) next(handler Handle, content string) bool {
	if jh.handle != nil {
		return jh.handle.Handle(content)
	}
	return false
}
