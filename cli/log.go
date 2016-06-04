package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func initLogger(c *cli.Context) error {
	level, err := log.ParseLevel(c.String("log-level"))
	if err != nil {
		return err
	}

	log.SetLevel(level)
	logLevel = level
	if c.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	switch strings.ToLower(c.String("log-formatter")) {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	}

	var logFile *os.File
	logFileName := c.String("log-file")
	if logFileName != "" {
		logFile, err = os.OpenFile(
			logFileName,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			os.FileMode(c.Int("log-file-mode")),
		)
		if err != nil {
			return err
		}
	} else {
		logFile = os.Stdout
	}
	log.SetOutput(logFile)

	return nil
}
