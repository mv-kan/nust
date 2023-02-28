package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mv-kan/nust/console"
	"github.com/mv-kan/nust/core"
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "undo a nust task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		makeargs, _ := cmd.Flags().GetString("makeargs")
		nocolor, _ := cmd.Flags().GetBool("no-color")
		if nocolor {
			color.NoColor = true // disables colorized output
		}

		var err error
		if force {
			err = core.UndoTaskForce(args[0], makeargs)
		} else {
			err = core.UndoTask(args[0], makeargs)
		}

		if err != nil {
			console.Danger(fmt.Sprintf("NUST ERROR: %v\n", err))

			os.Exit(1)
		}
	},
}

func init() {
	undoCmd.PersistentFlags().StringP("makeargs", "m", "", "pass flags and values to makefile")
	undoCmd.PersistentFlags().BoolP("force", "f", false, "run without checks in the exec info json file")
	// undoCmd.PersistentFlags().BoolP("no-auto-recovery", "n", true, "run without checks in the exec info json file")
	rootCmd.AddCommand(undoCmd)
}
