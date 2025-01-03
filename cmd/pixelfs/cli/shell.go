package cli

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/c-bata/go-prompt"
	"github.com/pixelfs/pixelfs/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(shellCmd)
}

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Switch to command-line mode",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		p := prompt.New(
			executor,
			completer,
			prompt.OptionLivePrefix(func() (prefix string, useLivePrefix bool) {
				pwd := viper.GetString("pwd")
				if pwd == "" {
					pwd = "/"
				}

				return pwd + " > ", true
			}),
			prompt.OptionAddKeyBind(prompt.KeyBind{
				Key: prompt.ControlC,
				Fn:  func(buf *prompt.Buffer) {},
			}),
		)

		log.ShellMode = true
		p.Run()
	},
}

func executor(input string) {
	input = strings.TrimSpace(input)
	if input == "exit" {
		os.Exit(0)
	}

	args := strings.Fields(input)
	if len(args) == 0 {
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}

			wg.Done()
		}()

		command, _, err := rootCmd.Find(args)
		if err != nil {
			log.Cli().Error().Err(err)
			return
		}

		command.Flags().VisitAll(func(flag *pflag.Flag) {
			flag.Value.Set(flag.DefValue)
		})

		rootCmd.SetArgs(args)
		if err := rootCmd.Execute(); err != nil {
			fmt.Println(err)
		}
	}()

	wg.Wait()
}

func completer(d prompt.Document) []prompt.Suggest {
	input := d.TextBeforeCursor()
	if input == "" || strings.HasSuffix(input, " ") {
		return nil
	}

	var suggestions []prompt.Suggest
	for _, cmd := range rootCmd.Commands() {
		suggestions = append(suggestions, prompt.Suggest{
			Text:        cmd.Name(),
			Description: cmd.Short,
		})
	}

	return prompt.FilterHasPrefix(suggestions, input, true)
}
