package main

import "sync"

//https://docs.google.com/document/d/1DJYHimjA-TjMF7eaEW7VBI6KJirWN9_TKkfYoraic5E/edit?tab=t.0

type Node[K comparable, V any] struct {
	key   K
	value V
	prev  *Node[K, V]
	next  *Node[K, V]
}

type LRUCache[K comparable, V any] struct {
	capacity int
	cache    map[K]*Node[K, V]
	head     *Node[K, V]
	tail     *Node[K, V]
	mu       *sync.Mutex
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	cache := &LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]*Node[K, V]),
		mu:       &sync.Mutex{},
	}
	cache.head = &Node[K, V]{}
	cache.tail = &Node[K, V]{}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (lr *LRUCache[K, V]) Get(key K) (V, bool) {
	lr.mu.Lock()
	defer lr.mu.Unlock()

	if node, exits := lr.cache[key]; exits {
		lr.moveToHead(node)
		return node.value, true
	}
	var zero V
	return zero, false
}

func (lr *LRUCache[K, V]) Put(key K, value V) {
	lr.mu.Lock()
	defer lr.mu.Unlock()

	if node, exists := lr.cache[key]; exists {
		node.value = value
		lr.moveToHead(node)
		return
	}

	newNode := &Node[K, V]{
		key:   key,
		value: value,
	}
	lr.cache[key] = newNode
	lr.moveToHead(newNode)

	if len(lr.cache) > lr.capacity {
		lastNode := lr.removeTail()
		delete(lr.cache, lastNode.key)
	}
}

func (lr *LRUCache[K, V]) addToHead(node *Node[K, V]) {
	node.next = lr.head.next
	node.prev = lr.head

	lr.head.next.prev = node
	lr.head.next = node
}

func (lr *LRUCache[K, V]) removeNode(node *Node[K, V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

}

func (lr *LRUCache[K, V]) moveToHead(node *Node[K, V]) {
	lr.removeNode(node)
	lr.addToHead(node)
}

func (lr *LRUCache[K, V]) removeTail() *Node[K, V] {
	node := lr.tail.prev
	lr.removeNode(node)
	return node
}

func (lr *LRUCache[K, V]) Size() int {
	lr.mu.Lock()
	defer lr.mu.Unlock()

	return len(lr.cache)
}
