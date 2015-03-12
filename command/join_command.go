// Copyright 2015 Axel Etcheverry.
//

package command

import (
    "os"
    "net"
    "fmt"
    "github.com/codegangsta/cli"
    "github.com/parnurzeal/gorequest"
    "github.com/fatih/color"
    "../config"
)

func NewJoinCommand() cli.Command {
    return cli.Command{
        Name:  "join",
        Usage: "Join the cluster",
        Action: joinCommandFunc,
    }
}

// joinCommandFunc executes the "join" command.
func joinCommandFunc(c *cli.Context) {

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

    var ip string = ""

    addrs, err := net.InterfaceAddrs()
    if err != nil {
        fmt.Printf("Oops: %v\n", err.Error())
        os.Exit(1)
    }

    for _, a := range addrs {
        if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                ip = ipnet.IP.String()
            }
        }
    }

    //fmt.Fprintf(color.Output, " %s Hostname: %s\n", color.GreenString("*"), name)
    //fmt.Fprintf(color.Output, " %s ip: %s\n", color.GreenString("*"), ip)

    resp, _, errs := gorequest.New().
        Put(config.Discovery + "/" + name).
        Send("value=" + ip).
        End()

    if (errs != nil) {
        fmt.Printf("Oops: %v\n", errs)
        os.Exit(1)
    }

    if (resp.StatusCode == 200 || resp.StatusCode == 201) {
        fmt.Fprintf(color.Output, " %s Join the cluster\t\t[ %s ]\n", color.GreenString("*"), color.GreenString("ok"))
    } else {
        fmt.Fprintf(color.Output, " %s Join the cluster\t\t[ %s ]\n", color.RedString("*"), color.RedString("!!"))
    }
}
