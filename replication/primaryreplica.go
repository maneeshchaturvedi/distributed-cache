package replication

import (
	"errors"

	"github.com/user/distributed-cache/cacheclient"
	"github.com/user/distributed-cache/cacheprotocol"
)

type PrimaryReplica struct{}

func NewPrimaryReplica() *PrimaryReplica {
	return &PrimaryReplica{}
}

func (pr *PrimaryReplica) PerformWriteOperation(client *cacheclient.CacheClient, key string, operation func(node *cacheclient.CacheNode) error) error {
	for i := 0; i < client.Replicas; i++ {
		nodeID, err := client.SortedMap.GetNode(key)
		if err != nil {
			return err
		}

		node, ok := client.Nodes[nodeID]
		if !ok {
			return errors.New("node not found")
		}

		err = operation(node)
		if err != nil {
			return err
		}

		key = nodeID
	}

	return nil
}

func (pr *PrimaryReplica) Get(client *cacheclient.CacheClient, key string) (interface{}, error) {
	nodeID, err := client.SortedMap.GetNode(key)
	if err != nil {
		return nil, err
	}

	node, ok := client.Nodes[nodeID]
	if !ok {
		return nil, errors.New("node not found")
	}

	return node.Get(key)
}
