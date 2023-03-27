package console

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Message(s string) {
	fmt.Println(s)
}

func Danger(s string) {
	c := color.New(color.FgRed, color.Bold)
	c.Fprintf(os.Stderr, "NUST: "+s)
}

func Warning(s string) {
	c := color.New(color.FgYellow, color.Bold)
	c.Print("NUST: " + s)
}
