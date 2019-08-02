package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	networkType string
)

func main() {
	app := cli.NewApp()
	app.Name = "troncli"
	app.Version = "0.0.1"
	app.Usage = "go version tron-cli"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "type, t",
			Usage:       "network type",
			Destination: &networkType,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "config",
			Action:  config,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
