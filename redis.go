package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

func getUserList() (string, error) {
	redis := NewRedisAi()
	get, err := redis.Get("user_list")
	if err != nil {
		return "", err
	}
	return get, nil
}

type RedisApi struct {
	redisPool   *redis.Pool
	redisServer string
}

func NewRedisAi() *RedisApi {
	pool := newRedisPoolWithSizeAndPasswd("127.0.0.1:6379", 10, "")
	return &RedisApi{
		redisPool:   pool,
		redisServer: "127.0.0.1:6379",
	}
}

func (api *RedisApi) Get(key string) (value string, err error) {
	// 获取一条Redis连接
	redisConn := api.redisPool.Get()
	r, err := redisConn.Do("GET", key)
	value, err = redis.String(r, err)
	redisConn.Close()
	return value, err
}

func newRedisPoolWithSizeAndPasswd(redisServer string, maxPoolSize int, passwd string) *redis.Pool {
	poolSize := 1000
	if maxPoolSize != 0 {
		poolSize = maxPoolSize
	}
	healthCheckPeriod := time.Second * 30
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 20,
		// Maximum number of connections allocated by the pool at a given time.
		// When zero, there is no limit on the number of connections in the pool
		MaxActive: poolSize,
		// Close connections after remaining idle for this duration. If the value
		// is zero, then idle connections are not closed. Applications should set
		// the timeout to a value less than the server's timeout.
		IdleTimeout: 300 * time.Second,
		// If Wait is true and the pool is at the MaxActive limit, then Get() waits
		// for a connection to be returned to the pool before returning.
		Wait: false,
		// Dial is an application supplied function for creating and configuring a
		// connection.
		//
		// The connection returned from Dial must not be in a special state
		// (subscribed to pub sub channel, transaction started, ...).
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisServer,
				redis.DialConnectTimeout(3*time.Second),
				// Read timeout on server should be greater than ping period.
				redis.DialReadTimeout(healthCheckPeriod+10*time.Second),
				redis.DialWriteTimeout(10*time.Second))
			if err != nil {
				return nil, err
			}
			// 密码非空才认证
			if passwd != "" {
				if _, err := c.Do("AUTH", passwd); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		// TestOnBorrow is an optional application supplied function for checking
		// the health of an idle connection before the connection is used again by
		// the application. Argument t is the time that the connection was returned
		// to the pool. If the function returns an error, then the connection is
		// closed.
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < healthCheckPeriod {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
