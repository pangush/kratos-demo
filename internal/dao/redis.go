package dao

import (
	"context"
	"kratos-demo/internal/conf"

	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/log"
)

func NewRedis() (r *redis.Redis, err error) {
	//var cfg struct {
	//	Client *redis.Config
	//}
	//if err = paladin.Get("redis.toml").UnmarshalTOML(&cfg); err != nil {
	//	return
	//}
	//r = redis.NewRedis(cfg.Client)
	r = redis.NewRedis(conf.Conf.Client)
	return
}

func (d *dao) PingRedis(ctx context.Context) (err error) {
	if _, err = d.redis.Do(ctx, "SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}