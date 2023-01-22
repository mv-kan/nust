package core

import (
	"errors"
	"os"
	"path/filepath"
)

const NustPackExt = ".nustpack"

func NewNust(make Make) Nust {
	return nust{make: make}
}

type Nust interface {
	// args is everything after -m parameter
	DoTask(task Task, args ...string) error
	UndoTask(task Task, args ...string) error

	// seek file with .nustpack extension and run it as usual task
	DoPackage(dir string, args ...string) error
	UndoPackage(dir string, args ...string) error
}

type nust struct {
	make Make
}

func (n nust) DoTask(task Task, args ...string) error {
	return n.make.NustDo(task, args...)
}

func (n nust) UndoTask(task Task, args ...string) error {
	return n.make.NustUndo(task, args...)
}

func findExt(dir string, extension string) (string, error) {
	// read all files from directory
	files, err := os.ReadDir(dir)

	// error handling
	if errors.Is(os.ErrNotExist, err) {
		return "", ErrNotExist
	} else if errors.Is(os.ErrPermission, err) {
		return "", ErrPermission
	} else if err != nil {
		return "", err
	}

	// find file with NustPackExt
	var pack string
	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext == extension {
			pack = filepath.Join(dir, file.Name())
		}
	}
	if pack == "" {
		return "", ErrNotExist
	}
	return pack, nil
}

func (n nust) DoPackage(dir string, args ...string) error {
	// find package file with extension NustPackExt
	packpath, err := findExt(dir, NustPackExt)
	if err != nil {
		return err
	}
	// pass file to make
	return n.make.NustDo(NewTask(packpath), args...)
}

func (n nust) UndoPackage(dir string, args ...string) error {
	// find package file with extension NustPackExt
	packpath, err := findExt(dir, NustPackExt)
	if err != nil {
		return err
	}
	// pass file to make
	return n.make.NustUndo(NewTask(packpath), args...)
}
