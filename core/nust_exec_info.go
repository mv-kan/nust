package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
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
		log.Print(err)
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
