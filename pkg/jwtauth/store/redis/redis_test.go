package redis

import (
	"context"
	"github.com/bilibili/kratos/pkg/container/pool"
	xtime "github.com/bilibili/kratos/pkg/time"
	"testing"
	"time"

	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/stretchr/testify/assert"
)

const (
	addr = "127.0.0.1:6379"
)

func TestStore(t *testing.T) {
	c := &pool.Config{
		Active:      10,
		Idle:        10,
		IdleTimeout: xtime.Duration(10*time.Second),
	}
	cfg := &redis.Config{
		Config:c,
		Proto:"tcp",
		Name: "reds1", Addr: "127.0.0.1:6379",
		DialTimeout: xtime.Duration(10*time.Second),
		ReadTimeout: xtime.Duration(1*time.Second),
		WriteTimeout: xtime.Duration(1*time.Second),
	}
	r := redis.NewRedis(cfg)

	store := NewStore(r)

	key := "test1"
	ctx := context.Background()
	err := store.Invalidate(ctx, key, 7200)
	assert.Nil(t, err)

	b, err := store.IsInvalidate(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, true, b)
}
