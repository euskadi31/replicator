// Copyright 2015 Axel Etcheverry.
//

package config

import (
    "os"
    "fmt"
    "encoding/json"
    "io/ioutil"
)

func Open(path string) (config Configuration, err error) {
    var data Configuration

    b, err := ioutil.ReadFile(path)
    if err == nil {
        error := json.Unmarshal(b, &data)

        if (error != nil) {
            fmt.Printf("Oops: %v\n", error)
            os.Exit(1)
        }
    }

    return data, err
}
