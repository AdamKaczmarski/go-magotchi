package cmd

import (
    "gomagotchi/src"
	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get your Gomagotchi's statistics",
	Run: func(cmd *cobra.Command, args []string) {
        src.ActOnCommand(src.STATS)
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
