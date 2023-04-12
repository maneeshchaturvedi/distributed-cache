package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"time"

	"github.com/user/distributed-cache/cacheclient"
	"github.com/user/distributed-cache/cacheprotocol"
	"github.com/user/distributed-cache/replication"
)

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	hashFunc := fnv.New32
	replicationStrategy := replication.NewChainReplication() // or NewPrimaryReplica()
	client := cacheclient.NewCacheClient(hashFunc, 2, replicationStrategy)

	client.AddNode("node1")
	client.AddNode("node2")
	client.AddNode("node3")

	key := randomString(10)
	value := randomString(5)

	err := client.Set(key, value)
	if err != nil {
		fmt.Printf("Error setting value: %v\n", err)
	}

	retVal, err := client.Get(key)
	if err != nil {
		fmt.Printf("Error getting value: %v\n", err)
	}

	if retVal == value {
		fmt.Printf("Value correctly retrieved: %s\n", retVal)
	} else {
		fmt.Printf("Value not correctly retrieved. Expected: %s, Got: %s\n", value, retVal)
	}
}
