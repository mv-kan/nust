package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mv-kan/nust/console"
	"github.com/mv-kan/nust/core"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "nust",
	Version: version,
	Short:   "NUST is Not Unique Setup Tool that can help you with setting up whatever you want",
	Long: `NUST is Not Unique Setup Tool used to ease pain with setting up stuff with nust packages and tasks.
See more info about the tool at https://github.com/mv-kan/nust`,
	Example: ``,
	Run: func(cmd *cobra.Command, args []string) {
		nocolor, _ := cmd.Flags().GetBool("no-color")
		force, _ := cmd.Flags().GetBool("force")
		retries, _ := cmd.Flags().GetInt("retries")
		sh, _ := cmd.Flags().GetString("nust-sh")

		os.Setenv(core.ENV_NUST_SH, sh)

		if nocolor {
			color.NoColor = true // disables colorized output
		}
		i := 0
		for {
			var err error

			filescript, args := args[0], args[1:]

			if force {
				err = core.NustRunForce(filescript, args...)
			} else {
				err = core.NustRun(filescript, args...)
			}

			if err != nil {
				console.Danger(fmt.Sprintf("(\"%s\" try number %d): %v\n", filescript, i, err))
				i++
				if i >= retries {
					os.Exit(1)
					break
				}
			} else {
				break
			}
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().Bool("no-color", false, "no color in output")
	rootCmd.PersistentFlags().BoolP("force", "f", false, "run without checks in the exec info json file")
	rootCmd.PersistentFlags().IntP("retries", "r", 0, "no more tries for this do task command")
	rootCmd.PersistentFlags().String("nust-sh", "/bin/bash", "script executor, but you can set it to /bin/sh like this: --nust-sh=\"/bin/sh\"")
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
