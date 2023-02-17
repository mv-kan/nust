package core

import (
	"os"
	"os/exec"
)

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func runTarget(file string, target string, makeargs ...string) error {
	makeargs = append([]string{target}, makeargs...)
	makeargs = append([]string{"-f", file}, makeargs...)

	// if you got empty string in makeargs it will break exec.Command func
	makeargs = deleteEmpty(makeargs)

	cmd := exec.Command("make", makeargs...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}

func nustDo(file string, makeargs ...string) error {
	return runTarget(file, "nust_do", makeargs...)
}

func nustUndo(file string, makeargs ...string) error {
	return runTarget(file, "nust_undo", makeargs...)
}
