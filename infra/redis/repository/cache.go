package repository

import (
	"context"
	"errors"
	"time"

	"base-gin-golang/domain/repository"
	errorConstants "base-gin-golang/errors"
	redisDb "base-gin-golang/infra/redis"

	redis "github.com/go-redis/redis/v8"
)

type cacheRepository struct {
	db *redisDb.Database
}

func NewCacheRepository(db *redisDb.Database) repository.CacheRepository {
	return &cacheRepository{
		db: db,
	}
}

func (r *cacheRepository) Get(ctx context.Context, key string) (string, error) {
	data, err := r.db.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", errorConstants.ErrKeyDoesNotExist
	}
	if err != nil {
		return "", err
	}
	return data, nil
}

func (r *cacheRepository) Set(
	ctx context.Context,
	key string,
	value interface{},
	timeToLive time.Duration,
) error {
	err := r.db.Set(ctx, key, value, timeToLive).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *cacheRepository) Del(ctx context.Context, key string) error {
	err := r.db.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *cacheRepository) Incr(ctx context.Context, key string) error {
	err := r.db.Incr(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
