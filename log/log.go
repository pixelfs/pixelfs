package log

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pixelfs/pixelfs/util"
	"github.com/jagottsicher/termcolor"
	arpc "github.com/lesismal/arpc/log"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	ShellMode = false
	Logger    = log.Logger
)

func SetLoggerColors() {
	var colors bool
	switch l := termcolor.SupportLevel(os.Stderr); l {
	case termcolor.Level16M:
		colors = true
	case termcolor.Level256:
		colors = true
	case termcolor.LevelBasic:
		colors = true
	case termcolor.LevelNone:
		colors = false
	default:
		colors = false
	}

	if _, noColorIsSet := os.LookupEnv("NO_COLOR"); noColorIsSet {
		colors = false
	}

	zerolog.ErrorMarshalFunc = func(err error) any {
		return strings.TrimPrefix(err.Error(), "internal: ")
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	home, _ := util.GetHomeDir()
	logDir := filepath.Join(home, "logs")
	if err := util.EnsureDir(logDir); err != nil {
		log.Fatal().Err(err).Msg("failed to create log directory")
	}

	logFile, err := os.OpenFile(filepath.Join(logDir, "pixelfs.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open log file")
	}

	Logger = log.Output(zerolog.MultiLevelWriter(
		zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: !colors},
		zerolog.ConsoleWriter{Out: logFile, TimeFormat: time.RFC3339, NoColor: true},
	))
}

func Debug() *zerolog.Event {
	return Logger.Debug()
}

func Info() *zerolog.Event {
	return Logger.Info()
}

func Warn() *zerolog.Event {
	return Logger.Warn()
}

func Error() *zerolog.Event {
	return Logger.Error()
}

func Fatal() *zerolog.Event {
	return Logger.Fatal()
}

type ArpcLogger struct {
	arpc.Logger
}

func (a *ArpcLogger) log(level zerolog.Level, format string, v ...any) {
	fields := strings.Split(fmt.Sprintf(format, v...), "\t")
	if len(fields) < 3 {
		return
	}

	tag := fields[0]
	addr := fields[1]
	message := strings.ToLower(fields[2])

	entry := Logger.With().Str("addr", addr).Logger()

	switch level {
	case zerolog.DebugLevel:
		entry.Debug().Msgf("%v %v", tag, message)
	case zerolog.InfoLevel:
		entry.Info().Msgf("%v %v", tag, message)
	case zerolog.WarnLevel:
		entry.Warn().Msgf("%v %v", tag, message)
	case zerolog.ErrorLevel:
		entry.Error().Msgf("%v %v", tag, message)
	default:
		entry.Info().Msgf("%v %v", tag, message)
	}
}

func (a *ArpcLogger) Debug(format string, v ...any) {
	a.log(zerolog.DebugLevel, format, v...)
}

func (a *ArpcLogger) Info(format string, v ...any) {
	a.log(zerolog.InfoLevel, format, v...)
}

func (a *ArpcLogger) Warn(format string, v ...any) {
	a.log(zerolog.WarnLevel, format, v...)
}

func (a *ArpcLogger) Error(format string, v ...any) {
	a.log(zerolog.ErrorLevel, format, v...)
}

type CliLogger struct {
	level zerolog.Level
}

func Cli() *CliLogger {
	return &CliLogger{level: zerolog.InfoLevel}
}

func (c *CliLogger) Error() *CliLogger {
	c.level = zerolog.ErrorLevel
	return c
}

func (c *CliLogger) Fatal() *CliLogger {
	c.level = zerolog.FatalLevel
	return c
}

func (c *CliLogger) Msg(a any) {
	if c.level != zerolog.FatalLevel {
		fmt.Println(a)
		return
	}

	if ShellMode {
		panic(a)
	} else {
		fmt.Println(a)
		os.Exit(1)
	}
}

func (c *CliLogger) Err(err error) {
	c.Msg(strings.TrimPrefix(err.Error(), "internal: "))
}

func (c *CliLogger) Msgf(format string, a ...any) {
	c.Msg(fmt.Sprintf(format, a...))
}
