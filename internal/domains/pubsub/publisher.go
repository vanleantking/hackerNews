package pubsub

import (
	"context"
	"hackerNewsApi/internal/infrastructure/pubsub"
)

type PublisherBus interface {
	Publisher(ctx context.Context, event pubsub.PubSubEvent) error
}
