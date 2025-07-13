package main

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strconv"
	"sync"
)

type ConsistentHash struct {
	Replicas int
	HashRing []int
	HashMap  map[int]string
	Node     map[string]bool
	mu       sync.RWMutex
}

// NewConsistentHash Initializes an empty hash ring with the given number of virtual nodes (replicas)
func NewConsistentHash(replicas int) *ConsistentHash {
	return &ConsistentHash{
		Replicas: replicas,
		HashMap:  make(map[int]string),
		Node:     make(map[string]bool),
	}
}

// Converts strings to a consistent int hash using the first 4 bytes of SHA256.
// Ensures same input always hashes to the same value.
func (ch *ConsistentHash) hash(data string) int {
	hash := sha256.Sum256([]byte(data))
	return int((uint32(hash[0]) << 24) | (uint32(hash[1]) << 16) | (uint32(hash[2]) << 8) | uint32(hash[3]))
}

func (ch *ConsistentHash) AddNode(node string) {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	if ch.Node[node] {
		return
	}
	ch.Node[node] = true
	for i := 0; i < ch.Replicas; i++ {
		virtualNode := node + "#" + strconv.Itoa(i)
		h := ch.hash(virtualNode)
		ch.HashRing = append(ch.HashRing, h)
		ch.HashMap[h] = node
	}
	sort.Ints(ch.HashRing)
}

func (ch *ConsistentHash) RemoveNode(node string) {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	if !ch.Node[node] {
		return
	}
	delete(ch.Node, node) //remove the node
	var newRing []int
	for i := 0; i < ch.Replicas; i++ {
		virtualNode := node + "#" + strconv.Itoa(i)
		h := ch.hash(virtualNode)
		delete(ch.HashMap, h) //remove from hashMap
	}
	for _, h := range ch.HashRing {
		if ch.HashMap[h] != "" {
			newRing = append(newRing, h) //remove from hashRing
		}
	}
	ch.HashRing = newRing
}

func (ch *ConsistentHash) GetNode(key string) string {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	if len(ch.HashRing) == 0 {
		return ""
	}
	h := ch.hash(key)

	idx := sort.Search(len(ch.HashRing), func(i int) bool {
		return ch.HashRing[i] >= h
	})

	if idx == len(ch.HashRing) {
		idx = 0
	}
	return ch.HashMap[ch.HashRing[idx]]
}

func main() {
	fmt.Println("Welcome to consistent-hashing!")

	ch := NewConsistentHash(3)

	ch.AddNode("NodeA")
	ch.AddNode("NodeB")
	ch.AddNode("NodeC")

	keys := []string{"apple", "banana", "cherry", "date", "elderberry"}

	for _, key := range keys {
		fmt.Printf("Key '%s' is assigned to %s\n", key, ch.GetNode(key))
	}
	fmt.Println("----------------------------------------")

	ch.AddNode("NodeD")
	for _, key := range keys {
		fmt.Printf("Key '%s' is assigned to %s\n", key, ch.GetNode(key))
	}

	fmt.Println("---------------------------------------")

	ch.RemoveNode("NodeD")
	for _, key := range keys {
		fmt.Printf("Key '%s' is assigned to %s\n", key, ch.GetNode(key))
	}

	fmt.Println("---------------------------------------")

	ch.RemoveNode("NodeC")
	for _, key := range keys {
		fmt.Printf("Key '%s' is assigned to %s\n", key, ch.GetNode(key))
	}
}
