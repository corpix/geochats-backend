package cli

import (
	"github.com/corpix/go-boilerplate/cli/commands/greet"
	"github.com/urfave/cli"
)

var (
	commands = []cli.Command{
		greet.Command,
	}
)
