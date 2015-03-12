// Copyright 2015 Axel Etcheverry.
//

package command

import (
    "os"
    "fmt"
    "encoding/json"
    "github.com/codegangsta/cli"
    "github.com/parnurzeal/gorequest"
    "github.com/fatih/color"
    "../etcd"
    "../config"
)

func NewListCommand() cli.Command {
    return cli.Command{
        Name:  "list",
        Usage: "List all node on the cluster",
        Action: listCommandFunc,
    }
}

// listCommandFunc executes the "list" command.
func listCommandFunc(c *cli.Context) {

    config, err := config.Open(c.GlobalString("config"))

    if (err != nil) {
        fmt.Fprintf(color.Output, " %s %s\n", color.RedString("*"), err)
        os.Exit(1)
    }

    resp, body, errs := gorequest.New().
        Get(config.Discovery).
        End()

    if (errs != nil) {
        fmt.Printf("Oops: %v\n", errs)
        os.Exit(1)
    }

    if (resp.StatusCode == 200) {

        var data etcd.Response
        error := json.Unmarshal([]byte(body), &data)

        if (error != nil) {
            fmt.Printf("Oops: %v\n", error)
            os.Exit(1)
        }

        for _,element := range data.Node.Nodes {
            fmt.Fprintf(color.Output, " %s %s : %s\n", color.GreenString("*"), element.Key[49:], element.Value)
        }

    } else {
        fmt.Fprintf(color.Output, " %s Leaving the cluster\t\t[ %s ]\n", color.RedString("*"), color.RedString("!!"))
    }
}
