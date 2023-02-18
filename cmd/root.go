package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "nust",
	Version: version,
	Short:   "NUST is Not Unique Setup Tool that can help you with setting up whatever you want",
	Long: `
NUST is Not Unique Setup Tool used to ease pain with setting up stuff with nust packages and tasks.
See more info about the tool at https://github.com/mv-kan/nust`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("use flag \"--help\" for more information")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
