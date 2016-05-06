package goredis

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

func InitRedis(host, port, socket string, db int, password string) (client *redis.Pool) {
	var network, addr string
	if socket != "" {
		network = "unix"
		addr = socket
	} else if host != "" && port != "" {
		network = "tcp"
		addr = host + ":" + port
	} else {
		panic("redis configuration error")
	}

	client = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(network, addr, redis.DialDatabase(db), redis.DialPassword(password))
			if err != nil {
				return nil, err
			} else {
				return conn, nil
			}
		},
		IdleTimeout: 60 * time.Second,
	}

	return client
}
