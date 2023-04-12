package cacheclient

import (
	"errors"
	"fmt"
	"sort"
)

type HashFunc func([]byte) uint32

type SortedMap struct {
	hashFunc HashFunc
	keys     []uint32
	nodes    map[uint32]string
}

func NewSortedMap(hashFunc HashFunc) *SortedMap {
	return &SortedMap{
		hashFunc: hashFunc,
		keys:     []uint32{},
		nodes:    make(map[uint32]string),
	}
}

func (sm *SortedMap) AddNode(nodeID string) {
	hash := sm.hashFunc([]byte(nodeID))
	sm.keys = append(sm.keys, hash)
	sm.nodes[hash] = nodeID
	sort.Slice(sm.keys, func(i, j int) bool {
		return sm.keys[i] < sm.keys[j]
	})
}

func (sm *SortedMap) RemoveNode(nodeID string) {
	hash := sm.hashFunc([]byte(nodeID))
	index := sm.nodeIndex(hash)
	sm.keys = append(sm.keys[:index], sm.keys[index+1:]...)
	delete(sm.nodes, hash)
}

func (sm *SortedMap) nodeIndex(hash uint32) int {
	index := sort.Search(len(sm.keys), func(i int) bool {
		return sm.keys[i] >= hash
	})
	return index % len(sm.keys)
}

func (sm *SortedMap) GetNode(key string) (string, error) {
	if len(sm.keys) == 0 {
		return "", errors.New("no nodes added")
	}
	hash := sm.hashFunc([]byte(key))
	index := sm.nodeIndex(hash)
	return sm.nodes[sm.keys[index]], nil
}
