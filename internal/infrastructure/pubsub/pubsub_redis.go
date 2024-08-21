package pubsub

import (
	"context"
	"encoding/json"
	pubsubBus "hackerNewsApi/internal/domains/pubsub"
	"hackerNewsApi/pkg/redis"
	"log"
	"sync"
)

type redisPubsubClient struct {
	pubsubClient redis.RedisClient
}

type PubSubClient interface {
}

func NewRedisClient(redisClient redis.RedisClient) (PubSubClient, error) {
	return &redisPubsubClient{
		pubsubClient: redisClient,
	}, nil
}

func (bus *redisPubsubClient) Publish(ctx context.Context, event pubsubBus.PubSubEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return bus.pubsubClient.
		GetClient().
		Publish(ctx, event.GetTopicName(), data).
		Err()
}

func (bus *redisPubsubClient) Subscribe(
	ctx context.Context,
	topic string, handler func(pubsubBus.PubSubEvent) error,
) error {
	pubsub := bus.pubsubClient.GetClient().Subscribe(ctx, topic)
	defer pubsub.Close()

	ch := pubsub.Channel()
	for msg := range ch {
		var event pubsubBus.PubSubEvent
		if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
			log.Println("Error unmarshalling event:", err)
			continue
		}

		if err := handler(event); err != nil {
			return err
		}
	}
	return nil
}

/**
 * Subcripbe for list
 * map[topic]handler : handler func(data pubsubBus.PubSubEvent) error
 */
type HandlerFunc func(data pubsubBus.PubSubEvent) error

func (r *redisPubsubClient) Subscribes(ctx context.Context, handers map[string]HandlerFunc) error {
	var wg sync.WaitGroup

	for topic, handler := range handers {
		wg.Add(1)
		go r.subscribeToTopic(ctx, topic, handler, &wg)
	}

	wg.Wait()
	return nil
}

func (r *redisPubsubClient) subscribeToTopic(
	ctx context.Context,
	topic string,
	handler HandlerFunc,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	r.Subscribe(ctx, topic, handler)
}
