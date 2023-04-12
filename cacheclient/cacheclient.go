package cacheclient

import (
	"fmt"

	"github.com/user/distributed-cache/cacheprotocol"
)

type CacheClient struct {
	SortedMap          *SortedMap
	Nodes              map[string]*CacheNode
	Replicas           int
	ReplicationStrategy cacheprotocol.Replication
}

func NewCacheClient(hashFunc HashFunc, replicas int, replicationStrategy cacheprotocol.Replication) *CacheClient {
	return &CacheClient{
		SortedMap:          NewSortedMap(hashFunc),
		Nodes:              make(map[string]*CacheNode),
		Replicas:           replicas,
		ReplicationStrategy: replicationStrategy,
	}
}

func (cc *CacheClient) AddNode(nodeID string) {
	node := NewCacheNode(nodeID)
	cc.Nodes[nodeID] = node
	cc.SortedMap.AddNode(nodeID)
}

func (cc *CacheClient) RemoveNode(nodeID string) {
	delete(cc.Nodes, nodeID)
	cc.SortedMap.RemoveNode(nodeID)
}

func (cc *CacheClient) performWriteOperation(key string, operation func(node *CacheNode) error) error {
	return cc.ReplicationStrategy.PerformWriteOperation(cc, key, operation)
}

func (cc *CacheClient) Set(key string, value interface{}) error {
	return cc.performWriteOperation(key, func(node *CacheNode) error {
		node.Set(key, value)
		return nil
	})
}

func (cc *CacheClient) Delete(key string) error {
	return cc.performWriteOperation(key, func(node *CacheNode) error {
		node.Delete(key)
		return nil
	})
}

func (cc *CacheClient) Get(key string) (interface{}, error) {
	return cc.ReplicationStrategy.Get(cc, key)
}
