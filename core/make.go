package core

import (
	"bytes"
	"fmt"
	"os/exec"
)

func runTarget(file string, target string, makeargs ...string) error {
	makeargs = append(makeargs, "-f", file)
	makeargs = append([]string{target}, makeargs...)
	cmd := exec.Command("make", makeargs...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	fmt.Printf("%s", out.String())
	if err != nil {
		return err
	}
	return nil
}

func NustDo(file string, makeargs ...string) error {
	return runTarget(file, "nust_do", makeargs...)
}

func NustUndo(file string, makeargs ...string) error {
	return runTarget(file, "nust_undo", makeargs...)
}
