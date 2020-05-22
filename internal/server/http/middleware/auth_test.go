package middleware

import (
	"context"
	"fmt"
	"kratos-demo/pkg/jwtauth"
	"testing"
	"time"

	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/container/pool"
	xtime "github.com/bilibili/kratos/pkg/time"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	jwtauthConf := &jwtauth.Config{
		Algo:   "HS512",
		TTL:    7200,
		Secret: "1234567890abcdef",
	}

	redisConf := &redis.Config{
		Config: &pool.Config{
			Active:      10,
			Idle:        10,
			IdleTimeout: xtime.Duration(10 * time.Second),
		},
		Proto: "tcp",
		Name:  "reds1", Addr: "127.0.0.1:6379",
		DialTimeout:  xtime.Duration(10 * time.Second),
		ReadTimeout:  xtime.Duration(1 * time.Second),
		WriteTimeout: xtime.Duration(1 * time.Second),
	}
	_, err := NewAuth(jwtauthConf, redisConf)

	if err != nil {
		t.Error(err)
	}
	token, err := AuthMiddleware.GenerateToken(context.Background(), "1")
	assert.Nil(t, err)

	fmt.Println(token)

	userID, err := AuthMiddleware.ParseUserID(context.Background(), token)
	assert.Nil(t, err)

	fmt.Println(userID)

}
