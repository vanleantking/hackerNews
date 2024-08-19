package pubsub

import (
	"context"
	"encoding/json"
	"hackerNewsApi/pkg/config"
	"hackerNewsApi/internal/domains/events"
	"log"

	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	client *redis.Client
}

type RedisClient interface {
}

func NewRedisClient(cfg config.Config) (RedisClient, error) {
	opt, err := redis.ParseURL(cfg.RedisURI)

	if err != nil {
		return nil, err
	}

	opt.PoolSize = cfg.RedisMaxActive
	opt.MinIdleConns = cfg.RedisMaxIdle

	client := redis.NewClient(opt)

	// Ping to test Redis connection
	if err = client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &redisClient{
		client: client,
	}, nil
}

func (bus *redisClient) Publish(ctx context.Context, event events.PubSubEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return bus.client.Publish(ctx, event.GetName(), data).Err()
}

func (bus *redisClient) Subscribe(
	ctx context.Context,
	eventName string, handler func(events.PubSubEvent),
) error {
	pubsub := bus.client.Subscribe(ctx, eventName)
	defer pubsub.Close()

	ch := pubsub.Channel()
	for msg := range ch {
		var event events.PubSubEvent
		if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
			log.Println("Error unmarshalling event:", err)
			continue
		}
		handler(event)
	}

	return nil
}
