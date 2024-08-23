package subscriber

import (
	"context"
	"fmt"
	"hackerNewsApi/internal/domains/pubsub"
)

type subscribeProcessing struct {
	Subscribe pubsub.RedisSubscribe
}

type SubscribeProcessing interface {
	Subscribes(ctx context.Context, handers map[string]pubsub.HandlerFunc) error
}

func NewSubscriberHandler(sub pubsub.RedisSubscribe) SubscribeProcessing {
	return &subscribeProcessing{
		Subscribe: sub,
	}
}

func (subscriberHdl *subscribeProcessing) Subscribes(
	ctx context.Context,
	handers map[string]pubsub.HandlerFunc,
) error {
	fmt.Println("enter Subscribes, ")
	return subscriberHdl.Subscribe.Subscribes(ctx, handers)
}
