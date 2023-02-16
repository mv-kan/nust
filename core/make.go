package core

import (
	"os"
	"os/exec"
)

func runTarget(file string, target string, makeargs ...string) error {
	makeargs = append(makeargs, "-f", file)
	makeargs = append([]string{target}, makeargs...)
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
