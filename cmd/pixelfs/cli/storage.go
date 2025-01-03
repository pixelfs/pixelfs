package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(storageCmd)
}

var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Storage management",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
