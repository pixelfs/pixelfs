package cli

import (
	"fmt"
	"os"

	"github.com/pixelfs/pixelfs/config"
	"github.com/pixelfs/pixelfs/log"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var cfgFile string

func init() {
	if len(os.Args) > 1 && (os.Args[1] == "version") {
		return
	}

	log.SetLoggerColors()
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().
		StringVarP(&cfgFile, "config", "c", "", "config file (default is ~/.pixelfs/config.toml)")
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = os.Getenv("PIXELFS_CONFIG")
	}
	if cfgFile != "" {
		err := config.LoadConfig(cfgFile, true)
		if err != nil {
			log.Fatal().Err(err).Msgf("error loading config file %s", cfgFile)
		}
	} else {
		err := config.LoadConfig("", false)
		if err != nil {
			log.Fatal().Err(err).Msgf("error loading config")
		}
	}

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get pixelfs configuration")
	}

	logLevel := zerolog.InfoLevel
	if cfg.Debug {
		logLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(logLevel)
}

var rootCmd = &cobra.Command{
	Use: "pixelfs",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
