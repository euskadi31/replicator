package main

import (
    "os"
    "github.com/codegangsta/cli"
    "./command"
)

func main() {
    app := cli.NewApp()
    app.Name = "replicator"
    app.Version = "0.0.1"
    app.Usage = "Replication on cluster"
    app.Flags = []cli.Flag{
        cli.BoolFlag{Name: "quiet", Usage: ""},
        cli.StringFlag{Name: "config", Value: "/etc/replicator.json", Usage: "Path to config file"},
    }
    app.Commands = []cli.Command{
        command.NewJoinCommand(),
        command.NewLeaveCommand(),
        command.NewListCommand(),
        command.NewSyncCommand(),
    }

    app.Run(os.Args)
}
