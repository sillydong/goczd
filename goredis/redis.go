package goredis

import (
	redigo "github.com/garyburd/redigo/redis"
	goredis "gopkg.in/redis.v5"
	"time"
)

func InitRedigo(host, port, socket string, db int, password string) (client *redigo.Pool) {
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

	client = &redigo.Pool{
		Dial: func() (redigo.Conn, error) {
			conn, err := redigo.Dial(network, addr, redigo.DialDatabase(db), redigo.DialPassword(password))
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

func InitRedisv3(host, port, socket string, db int64, password string) (client *goredis.Client, err error) {
	if socket != "" {
		client = goredis.NewClient(&goredis.Options{
			Network:     "unix",
			Addr:        socket,
			Password:    password,
			DB:          db,
			MaxRetries:  2,
			IdleTimeout: 60 * time.Second,
		})
	} else if host != "" && port != "" {
		client = goredis.NewClient(&goredis.Options{
			Network:     "tcp",
			Addr:        host + ":" + port,
			Password:    password,
			DB:          db,
			MaxRetries:  2,
			IdleTimeout: 60 * time.Second,
		})
	} else {
		panic("redis configuration error")
	}

	_, err = client.Ping().Result()

	return
}
