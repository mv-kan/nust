package console

import (
	"os"

	"github.com/fatih/color"
)

func Danger(s string) {
	c := color.New(color.FgRed, color.Bold)
	c.Fprintf(os.Stderr, "NUST: "+s)
}

func Warning(s string) {
	c := color.New(color.FgYellow, color.Bold)
	c.Print("NUST: " + s)
}
