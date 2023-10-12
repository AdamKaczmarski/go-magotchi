package cmd

import (
	"gomagotchi/internal"
	"github.com/spf13/cobra"
)

// petCmd represents the pet command
var petCmd = &cobra.Command{
	Use:   "pet",
	Short: "Pet you Gomagotchi",
	Run: func(cmd *cobra.Command, args []string) {
        internal.ActOnCommand(internal.PET)
	},
}

func init() {
	rootCmd.AddCommand(petCmd)
}
