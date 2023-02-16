package main

import (
	"fmt"

	"github.com/mv-kan/nust/core"
)

func main() {
	err := core.DoTask("test.nust", "SOME=test1")
	if err != nil {
		err = fmt.Errorf("NUST says there is an ERROR: %w", err)
		fmt.Println(err)
	}
	err = core.UndoTask("test.nust", "SOME=test2")
	if err != nil {
		err = fmt.Errorf("NUST says there is an ERROR: %w", err)
		fmt.Println(err)
	}
	fmt.Println("NUST HAPPY END")
}
