package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: actionSubCommands,
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "url",
			Value:    "",
			Usage:    "Base URL for the Kubermatic API",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "token",
			Value:    "",
			Usage:    "Bearer token for authentication with the Kubermatic API",
			Required: true,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
