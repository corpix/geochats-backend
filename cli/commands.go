package cli

import (
	"github.com/corpix/geochats-backend/cli/commands/greet"
	"github.com/urfave/cli"
)

var (
	commands = []cli.Command{
		greet.Command,
	}
)
