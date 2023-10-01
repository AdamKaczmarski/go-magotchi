package cmd

import (
	"gomagotchi/src"
	"github.com/spf13/cobra"
)

// petCmd represents the pet command
var petCmd = &cobra.Command{
	Use:   "pet",
	Short: "Pet you Gomagotchi",
	Run: func(cmd *cobra.Command, args []string) {
        src.ActOnCommand(src.PET)
	},
}

func init() {
	rootCmd.AddCommand(petCmd)
}
