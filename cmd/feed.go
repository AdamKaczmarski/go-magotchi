package cmd

import (
	"github.com/spf13/cobra"
	"gomagotchi/src"
)

// feedCmd represents the feed command
var feedCmd = &cobra.Command{
	Use:   "feed [burger || cake]",
	Short: "Feed your Gomagotchi with a burger or cake",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		//I don't like it but it works for now
		if args[0] == cmd.ValidArgs[0] {
			src.ActOnCommand(src.FEED_BURGER)
		} else {
			src.ActOnCommand(src.FEED_CAKE)
		}
	},
	ValidArgs: []string{"burger", "cake"},
}

func init() {
	rootCmd.AddCommand(feedCmd)
}
