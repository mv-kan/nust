package console

import "github.com/fatih/color"

func Danger(s string) {
	c := color.New(color.FgRed, color.Bold)
	c.Println(s)
}

func Warning(s string) {
	c := color.New(color.FgYellow, color.Bold)
	c.Println(s)
}
