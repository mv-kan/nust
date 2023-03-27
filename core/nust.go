package core

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mv-kan/nust/console"
)

func NustRun(file string, args ...string) error {
	// check if exists
	_, err := os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		return err
	}
	nustTask := NustTaskJSON{Filepath: file, Success: false}

	nustTasks := loadExecInfo(filepath.Dir(file))

	for _, task := range nustTasks {
		// if exists then don't run make target
		if task.Filepath == nustTask.Filepath && task.Success {
			console.Warning(fmt.Sprintf("skipping the \"%s\" file, because it was already executed and saved to \"%s\" in the same dir\n", file, execInfoFileName))
			return nil
		}
	}
	args = append([]string{file}, args...)

	err = runSh(args...)

	// depending on error we set status to false if fail and true if ok
	if err != nil {
		nustTask.Success = false
	} else {
		nustTask.Success = true
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

func NustRunForce(file string, args ...string) error {
	args = append([]string{file}, args...)

	err := runSh(args...)
	return err
}
