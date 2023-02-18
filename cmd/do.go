package cmd

import (
	"fmt"
	"os"

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

		var err error

		if force {
			err = core.DoTaskForce(args[0], makeargs)
		} else {
			err = core.DoTask(args[0], makeargs)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "NUST ERROR: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	doCmd.PersistentFlags().StringP("makeargs", "m", "", "pass flags and values to makefile")
	doCmd.PersistentFlags().BoolP("force", "f", false, "run without checks in the exec info json file")
	rootCmd.AddCommand(doCmd)
}
