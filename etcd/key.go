// Copyright 2015 Axel Etcheverry.
//

package etcd

type Key struct {
    Key string `json:"key"`
    Value string `json:"value"`
    ModifiedIndex uint64 `json:"modifiedIndex"`
    CreatedIndex uint64 `json:"createdIndex"`
}
