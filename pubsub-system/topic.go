package main

import "sync"

type Topic struct {
	name       string
	subscibers map[Subscriber]struct{}
	mu         sync.Mutex
}

func NewTopic(name string) *Topic {
	return &Topic{
		name:       name,
		subscibers: make(map[Subscriber]struct{}),
	}
}

func (t *Topic) AddSubscriber(subscriber Subscriber) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.subscibers[subscriber] = struct{}{}
}

func (t *Topic) RemoveSubscriber(subscriber Subscriber) {
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.subscibers, subscriber)
}

func (t *Topic) Publish(message *Message) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for subscriber := range t.subscibers {
		subscriber.OnMessage(message)
	}
}
