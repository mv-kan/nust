package cmd

import (
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

		if force {
			core.UndoTaskForce(args[0], makeargs)
		} else {
			core.UndoTask(args[0], makeargs)
		}
	},
}

func init() {
	undoCmd.PersistentFlags().StringP("makeargs", "m", "", "pass flags and values to makefile")
	undoCmd.PersistentFlags().BoolP("force", "f", false, "run without checks in the exec info json file")
	rootCmd.AddCommand(undoCmd)
}
