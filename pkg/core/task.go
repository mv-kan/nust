package core

func NewTask(path string) Task {
	return task{path: path}
}

type Task interface {
	GetPath() string
}

type task struct {
	path string
}

func (t task) GetPath() string {
	return t.path
}
