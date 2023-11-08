package redislib

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

type RedisConf struct {
	Addr   string
	Passwd string
}

type Redis struct {
	*redis.Client
	redsync *redsync.Redsync
}

func NewRedis(conf *RedisConf) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Passwd,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &Redis{
		Client:  client,
		redsync: redsync.New(goredis.NewPool(client)),
	}, nil
}

func (redis *Redis) Lock(ctx context.Context, key string, expire time.Duration) (*redsync.Mutex, error) {
	mutex := redis.redsync.NewMutex(key, redsync.WithExpiry(expire))
	if err := mutex.LockContext(ctx); err != nil {
		return nil, err
	}
	return mutex, nil
}
