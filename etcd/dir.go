// Copyright 2015 Axel Etcheverry.
//

package etcd

type Dir struct {
    Key string `json:"key"`
    Dir bool `json:"dir"`
    Nodes[] Key `json:"nodes"`
    ModifiedIndex uint64 `json:"modifiedIndex"`
    CreatedIndex uint64 `json:"createdIndex"`
}
