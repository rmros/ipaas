package redis

import (
	"ipaas/pkg/tools/configz"
	"ipaas/pkg/tools/log"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func init() {
	if configz.MustBool("redis", "requiredPW", false) {
		client = redis.NewClient(
			&redis.Options{
				Addr:     configz.GetString("redis", "addr", "127.0.0.1:6379"),
				Password: configz.GetString("redis", "password", "root@123"),
				DB:       configz.MustInt("redis", "db", 0),
			})
	} else {
		client = redis.NewClient(
			&redis.Options{
				Addr: configz.GetString("redis", "addr", "127.0.0.1:6379"),
				DB:   configz.MustInt("redis", "db", 0),
			})
	}

	if err := client.Ping().Err(); err != nil {
		log.Critical("init redis client err: %v", err)
	}
}

// GetClient return redis client
func GetClient() *redis.Client {
	return client
}
