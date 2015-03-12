// Copyright 2015 Axel Etcheverry.
//

package config

type Command struct {
    Cmd string `json:"cmd"`
    Option string `json:"option"`
    Path string `json:"path"`
    User string `json:"user"`
    Excludes[] string `json:"excludes"`
    Delete bool `json:"delete"`
}
