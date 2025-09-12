package main

import "fmt"

type Publisher struct {
	topics map[*Topic]struct{}
}

func NewPublisher() *Publisher {
	return &Publisher{
		topics: make(map[*Topic]struct{}),
	}
}

func (p *Publisher) RegisterTopic(topic *Topic) {
	p.topics[topic] = struct{}{}
}

func (p *Publisher) Publish(topic *Topic, message *Message) {
	if _, exists := p.topics[topic]; !exists {
		fmt.Printf("This publisher can't publish to topic: %s\n", topic.name)
		return
	}
	topic.Publish(message)
}
