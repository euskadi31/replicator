// Copyright 2015 Axel Etcheverry.
//

package config

type Configuration struct {
    Discovery string `json:"discovery"`
    Commands map[string]Command `json:"commands"`
}
