package pubsub

import (
	"context"
)

type RedisSubscribe interface {
	Subscribe(ctx context.Context, topic string, handler func(data []byte) error) error
	Subscribes(ctx context.Context, handers map[string]HandlerFunc) error
}
