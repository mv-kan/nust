package core

import (
	"fmt"
)

var (
	ErrNotDonePrev = fmt.Errorf("cannot find the task in \"%s\" with status true, try to use force undo with flag -f", execInfoFileName)
)
