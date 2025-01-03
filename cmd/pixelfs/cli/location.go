package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(locationCmd)
}

var locationCmd = &cobra.Command{
	Use:   "location",
	Short: "Location management",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
