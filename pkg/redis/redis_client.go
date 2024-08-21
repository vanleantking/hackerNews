package redis

import (
	"context"
	"hackerNewsApi/pkg/config"
	appLog "hackerNewsApi/pkg/logger"
	"sync"

	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	RedisClient *redis.Client
}

type RedisClient interface {
	GetClient() *redis.Client
}

var (
	once          sync.Once
	redisInstance *redisClient
)

func NewRedisClient(conf *config.Config, log appLog.Logger) (RedisClient, error) {
	once.Do(func() {
		client, err := connect(conf.RedisURI, conf.RedisMaxActive, conf.RedisMaxIdle)
		if err != nil {
			return
		}
		redisInstance = &redisClient{RedisClient: client}
	})
	return redisInstance, nil
}

func (rdClient *redisClient) GetClient() *redis.Client {
	return rdClient.RedisClient
}

type redisConfig struct {
	Host            string
	Port            string
	Password        string
	MaxIdle         int
	MaxActive       int
	MaxIdleTimeOut  int
	MaxConnLifetime int
	Wait            bool
}
type redisOption func(*redisConfig)

func connect(redisURI string, maxActive, maxIdle int) (*redis.Client, error) {
	opt, err := redis.ParseURL(redisURI)

	if err != nil {
		return nil, err
	}

	opt.PoolSize = maxActive
	opt.MinIdleConns = maxIdle

	client := redis.NewClient(opt)

	// Ping to test Redis connection
	if err = client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return client, nil
}

func SetMaxIdle(conns int) redisOption {
	return func(c *redisConfig) {
		if conns > 0 {
			c.MaxIdle = conns
		}
	}
}

func SetMaxIdleTimeout(timeout int) redisOption {
	return func(c *redisConfig) {
		if timeout > 0 {
			c.MaxIdleTimeOut = timeout
		}
	}
}

func SetMaxActive(conns int) redisOption {
	return func(c *redisConfig) {
		if conns > 0 {
			c.MaxActive = conns
		}
	}
}

func SetConnMaxLifetime(timeout int) redisOption {
	return func(c *redisConfig) {
		if timeout > 0 {
			c.MaxConnLifetime = timeout
		}
	}
}
