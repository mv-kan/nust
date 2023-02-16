package core

type NustTask struct {
	Filepath string // not full but relative to the nustpackage
}

type NustTaskJSON struct {
	Filepath string
	Status   bool // true OK, false NOT OK
}
