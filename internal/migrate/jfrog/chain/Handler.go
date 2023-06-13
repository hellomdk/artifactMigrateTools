package chain

type Handle interface {
	Handle(content string) bool
	next(handler Handle, content string) bool
}
