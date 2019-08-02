package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/urfave/cli"
)

func config(c *cli.Context) error {
	// Prepare some data to insert into the template.
	type NodeConfig struct {
		Type string
	}
	nc := NodeConfig{"mainnet"}

	// Create a new template and parse the letter into it.
	// t := template.Must(template.New("config").Parse(ConfigTemplate))
	t, err := template.ParseFiles("config.tmpl")
	if err != nil {
		return err
	}

	// save to file
	f, err := os.Create("config.conf")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return err
	}

	// Execute the template for each nc.
	err = t.Execute(f, nc)
	if err != nil {
		return err
	}

	return nil
}
