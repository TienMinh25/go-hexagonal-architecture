package storageredis

import (
	"context"

	"github.com/TienMinh25/go-hexagonal-architecture/config"
	"github.com/redis/go-redis/v9"
)

// NewRedis creates a new instance of client
func NewRedis(ctx context.Context, config *config.Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
