package core

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func DoTask(file string, makeargs ...string) error {
	// check if exists
	_, err := os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		return err
	}
	nustTask := NustTaskJSON{Filepath: file, Status: false}

	nustTasks := loadExecInfo(filepath.Dir(file))
	fmt.Println("nustTasks: ", nustTasks)
	for _, task := range nustTasks {
		// if exists then don't run make target
		if task.Filepath == nustTask.Filepath && task.Status {
			fmt.Printf("NUST: skip the \"%s\" file, because it was already executed and saved to \"%s\" in the same dir\n", file, execInfoFileName)
			return nil
		}
	}

	err = nustDo(file, makeargs...)
	if err != nil {
		return err
	}
	// in nustDo command may be other nust tasks so we need to update nustTasks obj in case if some nusttasks were finished
	nustTasks = loadExecInfo(filepath.Dir(file))
	fmt.Println("nustTasks: ", nustTasks)

	nustTask.Status = true

	nustTasks = append(nustTasks, nustTask)
	err = saveExecInfo(filepath.Dir(file), nustTasks)

	return err
}

func UndoTask(file string, makeargs ...string) error {
	// check if exists
	_, err := os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		return err
	}

	nustTasks := loadExecInfo(filepath.Dir(file))
	for i, task := range nustTasks {
		// if exists and status is positive
		if task.Filepath == file && task.Status {
			err = nustUndo(file, makeargs...)
			_ = saveExecInfo(filepath.Dir(file), remove(nustTasks, i))

			return err
		}
	}
	return ErrNotDonePrev
}

func DoTaskForce(file string, makeargs ...string) error {
	err := nustDo(file, makeargs...)
	return err
}

func UndoTaskForce(file string, makeargs ...string) error {
	err := nustUndo(file, makeargs...)
	return err
}
