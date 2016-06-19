package serve

import (
	log "github.com/Sirupsen/logrus"
	"github.com/corpix/geochats-backend/api"
	"github.com/corpix/geochats-backend/config"
	"github.com/urfave/cli"
	"net/http"
	"time"
)

var (
	// Command describes the command of some level
	// to be used in the top-level application
	Command = cli.Command{
		Name:      "serve",
		ShortName: "g",
		Usage:     "serves the API",
		Flags:     Flags,
		Action:    ServeAction,
	}

	// Flags contains all available command-line action flags
	Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "database-addr, d",
			Value:  "mongodb://127.0.0.1:27017",
			EnvVar: "DATABASE_ADDR",
			Usage:  "database addr to connect to",
		},
		cli.StringFlag{
			Name:   "database-name, n",
			Value:  "geochats",
			EnvVar: "DATABASE_NAME",
			Usage:  "database name to store data in",
		},
		cli.StringFlag{
			Name:   "listen-addr, l",
			Value:  ":3000",
			EnvVar: "LISTEN_ADDR",
			Usage:  "addr to serve on(eg 127.0.0.1:3000)",
		},
	}
)

// ServeAction is an entry point for the serve command of the API server
func ServeAction(ctx *cli.Context) error {
	conf := &config.Config{
		ListenAddr:             ctx.String("listen-addr"),
		DatabaseAddr:           ctx.String("database-addr"),
		DatabaseName:           ctx.String("database-name"),
		DatabaseConnectTimeout: 5 * time.Second,
	}
	config.Set(conf)

	for {
		router, err := api.New()
		if err != nil {
			return err
		}

		log.Infof("Serving on addr %s...", conf.ListenAddr)
		if err := http.ListenAndServe(conf.ListenAddr, router); err != nil {
			log.Errorf("Failed to accept new connection: %+v", err)
			time.Sleep(1 * time.Second)
		}
	}
}
