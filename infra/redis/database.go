package redis

import (
	"base-gin-golang/config"
	"context"
	"crypto/tls"

	"github.com/go-redis/redis/v8"
)

type Database struct {
	*redis.Client
}

func ConnectRedis(cfg *config.Environment) (*Database, error) {
	var tlsConfig *tls.Config
	if cfg.RedisUseSSL {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true, //nolint:gosec // debug
		}
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:      cfg.RedisURI,
		Password:  cfg.RedisPassword,
		TLSConfig: tlsConfig,
	})
	err := rdb.Ping(context.Background()).Err()
	return &Database{rdb}, err
}
