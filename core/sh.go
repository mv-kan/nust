package core

import (
	"os"
	"os/exec"
)

// this file contains functionality for calling sh scripts or bash scripts depending on ENV VARIABLE
// called NUST_SH that contains path to /bin/sh or /bin/bash
// default value is /bin/bash
const (
	ENV_NUST_SH   = "NUST_SH"
	nustShDefault = "/bin/bash"
)

func runSh(args ...string) error {
	// sh is path to script executor
	sh := os.Getenv(ENV_NUST_SH)

	if sh == "" {
		sh = nustShDefault
	}

	cmd := exec.Command(sh, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}
