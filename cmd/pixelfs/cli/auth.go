package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(authCmd)
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with the pixelfs",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
