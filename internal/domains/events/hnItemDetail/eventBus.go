package hnitemdetail

import (
	"context"
	"hackerNewsApi/internal/domains/events"
)

type EventBus interface {
	Publish(ctx context.Context, event events.PubSubEvent) error
	Subscribe(ctx context.Context, eventName string, handler func(events.PubSubEvent)) error
}
