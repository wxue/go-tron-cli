package main

import (
    "os/exec"

    "github.com/urfave/cli"
)

func run(c *cli.Context) error {
    cmd := exec.Command("java",
        "-jar", "./FullNode.jar",
        "-c", "./config.conf")
    err := cmd.Run()
    if err != nil {
        return err
    }
    return nil
}
