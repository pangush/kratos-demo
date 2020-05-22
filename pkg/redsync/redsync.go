package redsync

import (
	"kratos-demo/pkg/redsync/redis"

	"github.com/go-redsync/redsync"
)

// NewRedsync returns a new redsync.Redsync
func NewRedsync(configs []redis.Config) *redsync.Redsync {
	pools := redis.NewPools(configs)
	return redsync.New(pools)
}
