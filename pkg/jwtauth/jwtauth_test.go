package jwtauth

import (
	"context"
	"fmt"
	redis2 "kratos-demo/pkg/jwtauth/store/redis"
	"testing"
	"time"

	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/container/pool"
	xtime "github.com/bilibili/kratos/pkg/time"
	"github.com/stretchr/testify/assert"
)

func jwt() (*Jwtauth, error) {
	c := &pool.Config{
		Active:      10,
		Idle:        10,
		IdleTimeout: xtime.Duration(10 * time.Second),
	}

	r := redis.NewRedis(&redis.Config{
		Config: c,
		Proto:  "tcp",
		Name:   "reds1", Addr: "127.0.0.1:6379",
		DialTimeout:  xtime.Duration(10 * time.Second),
		ReadTimeout:  xtime.Duration(1 * time.Second),
		WriteTimeout: xtime.Duration(1 * time.Second),
	})

	jwt, err := New(&Config{
		Algo:   "HS512",
		TTL:    7200,
		Secret: "1234567890abcdef",
	}, redis2.NewStore(r))

	return jwt, err
}

func TestJwt_GenerateToken(t *testing.T) {
	jwt, err := jwt()
	assert.Nil(t, err)

	userID := "1111"
	token, err := jwt.GenerateToken(userID)
	assert.Nil(t, err)
	fmt.Println(token)

	userID, err = jwt.GetUserID(context.Background(), token)
	assert.Nil(t, err)
	fmt.Println(userID)

	err = jwt.DestroyToken(context.Background(), token)
	assert.Nil(t, err)

	// userID, err = jwt.GetUserID(context.Background(), token)
	// assert.Nil(t, err)
	// fmt.Println(userID)
}
