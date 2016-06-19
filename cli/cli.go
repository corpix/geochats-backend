package cli

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"path"
	"syscall"
)

var (
	logLevel log.Level
	version  = ""
)

func appBefore(c *cli.Context) error {
	var err error
	for _, fn := range [](func(*cli.Context) error){initLogger} {
		err = fn(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func toggleDebugLevel() {
	currLevel := log.GetLevel()
	if currLevel == log.DebugLevel {
		log.Info("Disabling debug log level")
		log.SetLevel(logLevel)
	} else {
		log.Info("Enabling debug log level")
		log.SetLevel(log.DebugLevel)
	}
}

func signalingLoop(sigChan chan os.Signal) {
MainLoop:
	for {
		sig := <-sigChan
		switch sig {
		case syscall.SIGUSR1:
			toggleDebugLevel()
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			defer os.Exit(0)
			log.Warn("Exiting with 0 code")
			close(sigChan)
			break MainLoop
		case syscall.SIGCHLD:
			log.Debug("Got SIGCHLD")
		default:
			log.Debug("Got %+v, but don't know how to handle it", sig)
		}
	}
}

// Run function creates an Application instance and subscribing to the OS
// signals, then it returns the control
func Run() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = "APP_USAGE"
	app.Version = fmt.Sprintf("%s", version)
	app.Authors = []cli.Author{
		cli.Author{Name: "Dmitry Moskowski", Email: "me@corpix.ru"},
	}
	app.Flags = rootFlags
	app.Before = appBefore
	app.Commands = commands

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan)
	go signalingLoop(sigChan)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Got an error on initial run: %+v", err)
	}
}
