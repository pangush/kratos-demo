package lredsync

import (
	"kratos-demo/pkg/lredsync/lredis"

	"github.com/go-redsync/redsync"
)

// NewRedsync returns a new redsync.Redsync
func NewRedsync(configs []lredis.Config) *redsync.Redsync {
	pools := lredis.NewPools(configs)
	return redsync.New(pools)
}
