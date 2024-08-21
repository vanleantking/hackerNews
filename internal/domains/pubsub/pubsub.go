package pubsub

import (
	"context"
)

type PubSubEvent interface {
	GetTopicName() string
}

type PubSubEventBus interface {
	Publisher(ctx context.Context, event PubSubEvent) error
	Subscriber(ctx context.Context, eventName string, handler func(PubSubEvent)) error
}
