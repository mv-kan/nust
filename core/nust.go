package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

const execInfoFileName = "nust_exec_info.json"

func loadExecInfo(dir string) []NustTaskJSON {
	content, err := os.ReadFile(path.Join(dir, execInfoFileName))
	if errors.Is(err, os.ErrNotExist) {
		return []NustTaskJSON{}
	}
	if err != nil {
		log.Fatal(err)
	}

	nustTasks := []NustTaskJSON{}
	err = json.Unmarshal(content, &nustTasks)
	if err != nil {
		log.Fatal(err)
	}
	return nustTasks
}

func saveExecInfo(dir string, tasks []NustTaskJSON) error {
	// save nust task to json obj

	content, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(path.Join(dir, execInfoFileName), content, 0755) // write permission
	return err
}

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
			fmt.Printf("NUST: skip the \"%s\" file, because it was already executed and saved to \"%s\" in the same dir\n", file, execInfoFileName)
			return nil
		}
	}

	err = nustDo(file, makeargs...)
	if err != nil {
		return err
	}
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
	for _, task := range nustTasks {
		// if exists and status is positive
		if task.Filepath == file && task.Status {
			err = nustUndo(file, makeargs...)
			task.Status = false
			_ = saveExecInfo(filepath.Dir(file), nustTasks)

			return err
		}
	}
	return errors.New(fmt.Sprintf("cannot find \"%s\" in \"%s\" with status true, try to use force undo with flag -f", file, execInfoFileName))
}
