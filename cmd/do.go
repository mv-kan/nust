package cmd

import (
	"github.com/mv-kan/nust/core"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do a nust task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		makeargs, _ := cmd.Flags().GetString("makeargs")

		if force {
			core.DoTaskForce(args[0], makeargs)
		} else {
			core.DoTask(args[0], makeargs)
		}
	},
}

func init() {
	doCmd.PersistentFlags().StringP("makeargs", "m", "", "pass flags and values to makefile")
	doCmd.PersistentFlags().BoolP("force", "f", false, "run without checks in the exec info json file")
	rootCmd.AddCommand(doCmd)
}
