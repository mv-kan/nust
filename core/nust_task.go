package core

type NustTask struct {
	Name string
}

type NustTaskJSON struct {
	Name   string
	Status bool // true OK, false NOT OK
}
