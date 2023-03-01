package cmd

import (
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "undo a nust task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		doOrUndo(false, cmd, args)
	},
}

func init() {
	undoCmd.PersistentFlags().StringP("makeargs", "m", "", "pass flags and values to makefile")
	undoCmd.PersistentFlags().BoolP("force", "f", false, "run without checks in the exec info json file")
	undoCmd.PersistentFlags().Bool("no-more-tries", false, "no more tries for this undo task command")

	rootCmd.AddCommand(undoCmd)
}
