package greet

import (
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	// Subcommands lists supported command-line commands
	// for this module
	Subcommands = []cli.Command{
		cli.Command{
			Name:  "hello",
			Usage: "says hello",
			Action: func(cli *cli.Context) error {
				user := cli.GlobalString("user")
				log.Debugf("greeting the user %s with hello...", user)
				log.Infof("Hello, %s!", user)
				return nil
			},
		},
	}

	// Command describes the command of some level
	// to be used in the top-level application
	Command = cli.Command{
		Name:        "greet",
		ShortName:   "g",
		Usage:       "greets the user",
		Flags:       Flags,
		Subcommands: Subcommands,
	}

	// Flags contains all available command-line action flags
	Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "user, u",
			Value:  "you",
			EnvVar: "USER",
			Usage:  "greet specified user",
		},
	}
)
