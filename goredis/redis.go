package goredis

import (
	"fmt"
	"gopkg.in/redis.v3"
	"reflect"
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

func FlatMap(v interface{}) []interface{} {
	args := make([]interface{}, 0)
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Map:
		for _, k := range rv.MapKeys() {
			args = append(args, k.Interface(), rv.MapIndex(k).Interface())
		}
	}
	return args
}

func FlatMapString(v interface{}) []string {
	args := make([]string, 0)
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Map:
		for _, k := range rv.MapKeys() {
			if rv.MapIndex(k).IsNil() {
				args = append(args, k.String(), "")
			} else {
				args = append(args, k.String(), fmt.Sprintf("%v", rv.MapIndex(k).Interface()))
			}
		}
	}
	return args
}
