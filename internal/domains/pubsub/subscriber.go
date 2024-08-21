package pubsub

import (
	"context"
	"hackerNewsApi/internal/infrastructure/pubsub"
)

type SubsriberBus interface {
	Subscriber(ctx context.Context, eventName string, handler func(pubsub.PubSubEvent)) error
	Subscribes(ctx context.Context, handers map[string]pubsub.HandlerFunc) error
}
