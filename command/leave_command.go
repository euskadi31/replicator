// Copyright 2015 Axel Etcheverry.
//

package command

import (
    "os"
    "fmt"
    "github.com/codegangsta/cli"
    "github.com/parnurzeal/gorequest"
    "github.com/fatih/color"
    "../config"
)

func NewLeaveCommand() cli.Command {
    return cli.Command{
        Name:  "leave",
        Usage: "Leave the cluster",
        Action: leaveCommandFunc,
    }
}

// leaveCommandFunc executes the "leave" command.
func leaveCommandFunc(c *cli.Context) {

    config, err := config.Open(c.GlobalString("config"))

    if (err != nil) {
        fmt.Fprintf(color.Output, " %s %s\n", color.RedString("*"), err)
        os.Exit(1)
    }

    name, err := os.Hostname()
    if err != nil {
        fmt.Printf("Oops: %v\n", err)
        os.Exit(1)
    }

    resp, _, errs := gorequest.New().
        Delete(config.Discovery + "/" + name).
        End()

    if (errs != nil) {
        fmt.Printf("Oops: %v\n", errs)
        os.Exit(1)
    }

    if (resp.StatusCode == 200) {
        fmt.Fprintf(color.Output, " %s Leaving the cluster\t\t[ %s ]\n", color.GreenString("*"), color.GreenString("ok"))
    } else {
        fmt.Fprintf(color.Output, " %s Leaving the cluster\t\t[ %s ]\n", color.RedString("*"), color.RedString("!!"))
    }
}
