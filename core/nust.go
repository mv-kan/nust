package core

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mv-kan/nust/console"
)

func DoTask(file string, makeargs ...string) error {
	// check if exists
	_, err := os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		return err
	}
	nustTask := NustTaskJSON{Filepath: file, Status: false}

	nustTasks := loadExecInfo(filepath.Dir(file))

	for _, task := range nustTasks {
		// if exists then don't run make target
		if task.Filepath == nustTask.Filepath && task.Status {
			console.Warning(fmt.Sprintf("NUST: skip the \"%s\" file, because it was already executed and saved to \"%s\" in the same dir\n", file, execInfoFileName))
			return nil
		}
	}

	err = nustDo(file, makeargs...)

	// depending on error we set status to false if fail and true if ok
	if err != nil {
		nustTask.Status = false
	} else {
		nustTask.Status = true
	}
	// in nustDo command may be other nust tasks so we need to update nustTasks obj in case if some nusttasks were finished
	nustTasks = loadExecInfo(filepath.Dir(file))

	// unoptimized I know but I don't care
	// we check if in new nustTasks there is already some task with same filepath
	// if there is same task we just change value in nustTasks, otherwise we add to nustTasks new element nustTask
	if len(nustTasks) == 0 {
		nustTasks = append(nustTasks, nustTask)
	}
	for i, task := range nustTasks {
		// if exists then don't run make target
		if task.Filepath == nustTask.Filepath {
			nustTasks[i] = nustTask
			break
		}
		// on last iter and no break means that not found
		if i+1 == len(nustTasks) {
			nustTasks = append(nustTasks, nustTask)
		}
	}

	saveerr := saveExecInfo(filepath.Dir(file), nustTasks)
	if saveerr != nil {
		log.Fatal(err)
	}

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
