package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	networkType  string
	fullNodePort string
	solNodePort  string
)

func main() {
	app := cli.NewApp()
	app.Name = "troncli"
	app.Version = "0.0.1"
	app.Usage = "go version tron-cli"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "ntype, nt",
			Usage:       "network type",
			Value:       "mainnet",
			Destination: &networkType,
		},
		cli.StringFlag{
			Name:        "fport, fp",
			Usage:       "fullnode port",
			Value:       "8090",
			Destination: &fullNodePort,
		},
		cli.StringFlag{
			Name:        "sport, sp",
			Usage:       "solidity port",
			Value:       "8091",
			Destination: &solNodePort,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init",
			Action:  initialize,
		},
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "config",
			Action:  configure,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
