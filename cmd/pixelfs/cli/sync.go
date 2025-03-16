package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(syncCmd)
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync management",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
