// Copyright 2015 Axel Etcheverry.
//

package command

import (
    "os"
    "os/exec"
    "log"
    "strconv"
    //"bytes"
    "fmt"
    "net"
    "encoding/json"
    "github.com/codegangsta/cli"
    "github.com/parnurzeal/gorequest"
    "github.com/fatih/color"
    "github.com/nightlyone/lockfile"
    "../etcd"
    "../config"
)

func NewSyncCommand() cli.Command {
    return cli.Command{
        Name:  "sync",
        Usage: "Sync on all node on the cluster",
        Action: syncCommandFunc,
    }
}

// syncCommandFunc executes the "sync" command.
func syncCommandFunc(c *cli.Context) {

    lock, err := lockfile.New("/tmp/replicator_sync.lck")
    if err != nil {
        fmt.Println("Cannot init lock. reason: %v", err)
        panic(err)
    }
    err = lock.TryLock()

    if err != nil {
        fmt.Fprintf(color.Output, " %s The sync process already started.\n", color.RedString("*"))
        os.Exit(1)
    }

    defer lock.Unlock()

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

        for name, command := range config.Commands {
            execCommand(name, command, data.Node.Nodes, ip)
        }
    }
}

func buildCommand(command config.Command, ip string) *exec.Cmd {
    args := []string{command.Option, command.Path, command.User + "@" + ip + ":" + command.Path}

    for _, exclude := range command.Excludes {
        args = append(args, "--exclude=" + exclude)
    }

    if (command.Delete) {
        args = append(args, "--delete")
    }

    if (command.Timeout > 0) {
        args = append(args, "--timeout=" + strconv.Itoa(command.Timeout))
    }

    return exec.Command(command.Cmd, args...)
}

func execSync(name string, host string, cmd *exec.Cmd) {

    out, err := cmd.Output()
    if err != nil {
        fmt.Fprintf(color.Output, " %s Sync %s on %s\t\t[ %s ]\n", color.RedString("*"), name, host, color.RedString("!!"))
        log.Print(err)
    } else {
        fmt.Fprintf(color.Output, " %s Sync %s on %s\t\t[ %s ]\n", color.GreenString("*"), name, host, color.GreenString("ok"))
        log.Print(out)
    }
}

func execCommand(name string, command config.Command, nodes[] etcd.Key, ip string) {

    for _, element := range nodes {

        if (element.Value != ip) {
            execSync(name, element.Key[49:], buildCommand(command, element.Value))
        }
    }
}
