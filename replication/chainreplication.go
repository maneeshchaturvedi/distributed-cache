package replication

import (
	"errors"

	"github.com/user/distributed-cache/cacheclient"
	"github.com/user/distributed-cache/cacheprotocol"
)

type ChainReplication struct{}

func NewChainReplication() *ChainReplication {
	return &ChainReplication{}
}

func (cr *ChainReplication) PerformWriteOperation(client *cacheclient.CacheClient, key string, operation func(node *cacheclient.CacheNode) error) error {
	nodeID, err := client.SortedMap.GetNode(key)
	if err != nil {
		return err
	}

	for i := 0; i < client.Replicas; i++ {
		node, ok := client.Nodes[nodeID]
		if !ok {
			return errors.New("node not found")
		}

		err := operation(node)
		if err != nil {
			return err
		}

		nodeID, err = client.SortedMap.GetNode(nodeID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cr *ChainReplication) Get(client *cacheclient.CacheClient, key string) (interface{}, error) {
	nodeID, err := client.SortedMap.GetNode(key)
	if err != nil {
		return nil, err
	}

	for i := 0; i < client.Replicas-1; i++ {
		nodeID, err = client.SortedMap.GetNode(nodeID)
		if err != nil {
			return nil, err
		}
	}

	node, ok := client.Nodes[nodeID]
	if !ok {
		return nil, errors.New("node not found")
	}

	return node.Get(key)
}
