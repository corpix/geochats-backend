package cli

import (
	"github.com/corpix/geochats-backend/cli/commands/serve"
	"github.com/urfave/cli"
)

var (
	commands = []cli.Command{
		serve.Command,
	}
)
