package goredis

import (
	"fmt"
	"gopkg.in/redis.v3"
	"time"
)

func InitRedis(host, port, socket string, db int64, password string) (client *redis.Client, err error) {
	if socket != "" {
		client = redis.NewClient(&redis.Options{
			Network:     "unix",
			Addr:        socket,
			Password:    password,
			DB:          db,
			MaxRetries:  2,
			IdleTimeout: 60 * time.Second,
		})
	} else if host != "" && port != "" {
		client = redis.NewClient(&redis.Options{
			Network:     "tcp",
			Addr:        host + ":" + port,
			Password:    password,
			DB:          db,
			MaxRetries:  2,
			IdleTimeout: 60 * time.Second,
		})
	} else {
		err = fmt.Errorf("redis configuration error")
	}

	_, err = client.Ping().Result()

	return
}
