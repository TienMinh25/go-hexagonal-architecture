package redis

import (
	"context"
	"time"

	storageredis "github.com/TienMinh25/go-hexagonal-architecture/infrastructure/storage/redis"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/adapter/config"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/port"
	"github.com/redis/go-redis/v9"
)

/**
 * Redis implements port.CacheRepository interface
 * and provides an access to the redis library
 */
type redisCache struct {
	client *redis.Client
}

func NewRedis(ctx context.Context, config *config.Redis) (port.CacheRepository, error) {
	client, err := storageredis.NewRedis(ctx, config)

	if err != nil {
		return nil, err
	}

	return &redisCache{
		client: client,
	}, err
}

// Set stores the value in the redis database
func (r *redisCache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

// Get retrieves the value from the redis database
func (r *redisCache) Get(ctx context.Context, key string) ([]byte, error) {
	res, err := r.client.Get(ctx, key).Result()
	bytes := []byte(res)
	return bytes, err
}

// Delete removes the value from the redis database
func (r *redisCache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// DeleteByPrefix removes the value from the redis database with the given prefix
func (r *redisCache) DeleteByPrefix(ctx context.Context, prefix string) error {
	var cursor uint64
	var keys []string

	for {
		var err error
		keys, cursor, err = r.client.Scan(ctx, cursor, prefix, 100).Result()
		if err != nil {
			return err
		}

		for _, key := range keys {
			err := r.client.Del(ctx, key).Err()
			if err != nil {
				return err
			}
		}

		if cursor == 0 {
			break
		}
	}

	return nil
}

// Close closes the connection to the redis database
func (r *redisCache) Close() error {
	return r.client.Close()
}
