package redis

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/cache/redis"
)

type Rediser interface {
	Do(ctx context.Context, commandName string, args ...interface{}) (reply interface{}, err error)

	// Close closes connection pool
	Close() error
}

// NewStore 创建基于redis存储实例
func NewStore(redis Rediser) *Store {
	return &Store{
		redis:  redis,
		prefix: "keyPrefix",
	}
}

// Store redis存储
type Store struct {
	prefix string
	redis  Rediser
}

func (s *Store) wrapperKey(key string) string {
	return fmt.Sprintf("%s:jwt.black:%s", s.prefix, key)
}

// Invalidate a token (add it to the blacklist).
func (s *Store) Invalidate(ctx context.Context, jti string, ttl int64) error {
	_, err := s.redis.Do(ctx, "SET", s.wrapperKey(jti), "1", "EX", ttl)

	return err
}

// Check token is Invalidate
func (s *Store) IsInvalidate(ctx context.Context, jti string) (bool, error) {
	bo, err := redis.Bool(s.redis.Do(ctx, "EXISTS", s.wrapperKey(jti)))

	if err == redis.ErrNil {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	if !bo {
		return false, nil
	}
	return true, nil
}
