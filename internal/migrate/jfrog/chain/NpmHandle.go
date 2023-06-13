package chain

import (
	"strings"
)

type NpmHandle struct {
	handle Handle
}

func (nh *NpmHandle) Handle(content string) bool {
	if strings.Contains(content, ".npm") {
		return true
	}
	return nh.next(nh.handle, content)
}

func (nh *NpmHandle) next(handler Handle, content string) bool {
	if nh.handle != nil {
		return nh.handle.Handle(content)
	}
	return false
}
