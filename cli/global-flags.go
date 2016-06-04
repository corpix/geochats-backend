package cli

import (
	"github.com/urfave/cli"
)

var (
	rootFlags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug",
			EnvVar: "DEBUG",
			Usage:  "debug mode",
		},
		cli.StringFlag{
			Name:  "log-level, l",
			Value: "info",
			Usage: "log level(debug, info, warn, error, fatal, panic)",
		},
		cli.StringFlag{
			Name:  "log-file",
			Value: "",
			Usage: "log file to log entries to",
		},
		cli.IntFlag{
			Name:  "log-file-mode",
			Value: 0600,
			Usage: "octal mode for log file", //
		},
		cli.StringFlag{
			Name:  "log-formatter",
			Value: "",
			Usage: "log formatter to use(available: json or none)",
		},
		cli.StringFlag{
			Name:  "config",
			Value: "config.xml",
			Usage: "path to configuration file",
		},
	}
)
