package pubsub

import (
	"context"
	"fmt"
	"hackerNewsApi/internal/domains/pubsub"
	"hackerNewsApi/pkg/redis"
	"sync"
)

type redisPubSubClient struct {
	pubSubClient redis.RedisClient
}

func NewRedisClient(redisClient redis.RedisClient) (pubsub.RedisPublish, pubsub.RedisSubscribe) {
	r := &redisPubSubClient{
		pubSubClient: redisClient,
	}
	return r, r
}

func (bus *redisPubSubClient) Publish(ctx context.Context, topic string, data []byte) error {
	fmt.Println("Published, ", topic, string(data))
	return bus.pubSubClient.
		GetClient().
		Publish(ctx, topic, data).
		Err()
}

func (bus *redisPubSubClient) Subscribe(ctx context.Context, topic string, handler func(data []byte) error,
) error {
	pubsub := bus.pubSubClient.GetClient().Subscribe(ctx, topic)
	defer pubsub.Close()

	ch := pubsub.Channel()
	for msg := range ch {
		if err := handler([]byte(msg.Payload)); err != nil {
			return err
		}
	}
	return nil
}

func (r *redisPubSubClient) Subscribes(ctx context.Context, handers map[string]pubsub.HandlerFunc) error {
	var wg sync.WaitGroup

	for topic, handler := range handers {
		wg.Add(1)
		go r.subscribeToTopic(ctx, topic, handler, &wg)
	}

	wg.Wait()
	return nil
}

func (r *redisPubSubClient) subscribeToTopic(
	ctx context.Context,
	topic string,
	handler pubsub.HandlerFunc,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	r.Subscribe(ctx, topic, handler)
}
