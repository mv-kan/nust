package core

type Make interface {
	// calls target nust_do in makefile
	NustDo(task Task, args ...string) error

	// calls target nust_undo in makefile
	NustUndo(task Task, args ...string) error
}
