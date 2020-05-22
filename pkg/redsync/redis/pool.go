package redis

import (
	"time"

	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
)

// Config client settings.
type Config struct {
	Addr string
	Auth string

	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration

	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// newPool returns a new redis.pool
func newPool(config Config) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Addr)
			if err != nil {
				return nil, err
			}
			if config.Auth != "" {
				if _, err := c.Do("AUTH", config.Auth); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// NewPools returns a slice redsync.Pool
func NewPools(configs []Config) []redsync.Pool {
	pools := make([]redsync.Pool, 0)
	for _, config := range configs {
		pool := newPool(config)
		pools = append(pools, pool)
	}

	return pools
}
