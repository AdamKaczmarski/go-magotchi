package cmd

import (
	"gomagotchi/internal"
	"github.com/spf13/cobra"
)

// gameCmd represents the game command
var gameCmd = &cobra.Command{
	Use:   "game",
    Short: "Play a game with your Gomagotchi",
	Long: "Play a guess if the number is higher or lower with your Gomagotchi to increase their happiness",
	Run: func(cmd *cobra.Command, args []string) {
        internal.ActOnCommand(internal.GAME)
	},
}

func init() {
	rootCmd.AddCommand(gameCmd)
}
