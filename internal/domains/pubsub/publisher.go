package pubsub

import (
	"context"
)

type RedisPublish interface {
	Publish(ctx context.Context, topic string, data []byte) error
}
