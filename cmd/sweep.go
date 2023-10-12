package cmd

import (
	"gomagotchi/internal"
	"github.com/spf13/cobra"
)

// sweepCmd represents the sweep command
var sweepCmd = &cobra.Command{
	Use:   "sweep",
	Short: "Sweep the poops",
	Run: func(cmd *cobra.Command, args []string) {
        internal.ActOnCommand(internal.SWEEP)
	},
}

func init() {
	rootCmd.AddCommand(sweepCmd)
}
