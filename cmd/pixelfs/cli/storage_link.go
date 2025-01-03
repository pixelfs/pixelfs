package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	storageCmd.AddCommand(storageLinkCmd)
}

var storageLinkCmd = &cobra.Command{
	Use:   "link",
	Short: "Storage link management",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
