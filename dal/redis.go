package dal

import (
	"context"
	"crud-apis-db-app/config"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type IRedisDb interface {
	Exists(ctx context.Context, key string) (bool, error)
	Keys(ctx context.Context, pattern string) ([]string, error)
	Create(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Read(ctx context.Context, key string) (string, error)
	Update(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	Ttl(ctx context.Context, key string) (time.Duration, error)
}

type Redis struct {
	Connection *redis.Client
}

var RedisDbClient = Redis{}

func NewRedis(config config.IConfig) (IRedisDb, error) {
	cfgRedis := config.Get().Redis

	options := &redis.Options{
		Addr:     cfgRedis.Address,
		Username: cfgRedis.Username,
		Password: cfgRedis.Password,
	}

	client := redis.NewClient(options)
	if client == nil {
		return nil, errors.New("error while initialising redis client")
	}

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	fmt.Println("Successfully connected to Redis")

	RedisDbClient = Redis{Connection: client}

	return &RedisDbClient, nil
}

func (r *Redis) Create(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	client := r.Connection

	err := client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	client := r.Connection

	err := client.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) Exists(ctx context.Context, key string) (bool, error) {
	client := r.Connection

	exists, err := client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if exists != 1 {
		return false, nil
	}

	return true, nil
}

func (r *Redis) Keys(ctx context.Context, pattern string) ([]string, error) {
	client := r.Connection
	keys, err := client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (r *Redis) Read(ctx context.Context, key string) (string, error) {
	client := r.Connection

	value, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}

func (r *Redis) Update(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	client := r.Connection

	err := client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) Ttl(ctx context.Context, key string) (time.Duration, error) {
	client := r.Connection
	ttl, err := client.TTL(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return ttl, nil
}
