package worker

type Task struct {
	Handler func(interface{})
	Args    interface{}
}
