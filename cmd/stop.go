package cmd

import (
	"gomagotchi/internal"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the Gomagotchi daemon",
	Run: func(cmd *cobra.Command, args []string) {
        internal.ActOnCommand(internal.STOP)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
