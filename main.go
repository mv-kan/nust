package main

import "github.com/mv-kan/nust/core"

func main() {
	core.NustDo("test.nust")
	core.NustUndo("test.nust")
}
