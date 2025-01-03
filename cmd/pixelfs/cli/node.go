package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nodeCmd)
}

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Node management",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
