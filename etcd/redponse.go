// Copyright 2015 Axel Etcheverry.
//

package etcd

type Response struct {
    Action string `json:"action"`
    Node Dir `json:"node"`
}
