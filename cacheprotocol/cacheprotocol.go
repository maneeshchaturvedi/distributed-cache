package cacheprotocol

import (
	"github.com/user/distributed-cache/cacheclient"
)

type Replication interface {
	PerformWriteOperation(client *cacheclient.CacheClient, key string, operation func(node *cacheclient.CacheNode) error) error
	Get(client *cacheclient.CacheClient, key string) (interface{}, error)
}
