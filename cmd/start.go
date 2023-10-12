package cmd

import (
    "gomagotchi/internal"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts and runs the Gomagotchi in the background.",
	Run: func(cmd *cobra.Command, args []string) {
        internal.ActOnCommand(internal.START)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
