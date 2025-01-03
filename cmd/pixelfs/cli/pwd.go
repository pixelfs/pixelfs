package cli

import (
	"fmt"

	"github.com/pixelfs/pixelfs/config"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pwdCmd)
}

var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Print working directory",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.GetConfig()
		if err != nil {
			log.Cli().Fatal().Err(err)
		}

		pwd := "/"
		if cfg.Pwd != "" {
			pwd = cfg.Pwd
		}

		fmt.Println(pwd)
	},
}
