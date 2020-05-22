package middleware

import (
	"context"
	"kratos-demo/pkg/bmcontext"
	"kratos-demo/pkg/jwtauth"
	storeRedis "kratos-demo/pkg/jwtauth/store/redis"

	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/ecode"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

// AuthMiddleware .
var AuthMiddleware *Auth

// Auth .
type Auth struct {
	Jwtauth *jwtauth.Jwtauth
}

// NewAuth returns a new Auth
func NewAuth(jwtauthConf *jwtauth.Config, redisConf *redis.Config) (*Auth, error) {
	redisClient := redis.NewRedis(redisConf)
	newJwtauth, err := jwtauth.New(jwtauthConf, storeRedis.NewStore(redisClient))

	if err != nil {
		return nil, err
	}

	AuthMiddleware = &Auth{Jwtauth: newJwtauth}

	return AuthMiddleware, nil
}

// AuthMiddleware 用户授权中间件
func (a *Auth) AuthMiddleware() bm.HandlerFunc {
	return func(c *bm.Context) {
		t := bmcontext.GetToken(c)
		if t == "" {
			c.JSON(nil, ecode.Unauthorized)
			c.AbortWithStatus(ecode.Unauthorized.Code())
			return
		}

		id, err := a.Jwtauth.ParseUserID(c, t)
		if err != nil {
			if err == jwtauth.ErrInvalidToken {
				c.JSON(nil, ecode.Unauthorized)
				c.AbortWithStatus(ecode.Unauthorized.Code())
				return
			}
			c.JSON(nil, ecode.Unauthorized)
			c.AbortWithStatus(ecode.Unauthorized.Code())
			return
		}

		if id == "" {
			c.JSON(nil, ecode.Unauthorized)
			c.AbortWithStatus(ecode.Unauthorized.Code())
			return
		}

		bmcontext.SetUserID(c, id)
		c.Next()
	}
}

// GenerateToken returns a single string of token
func (a *Auth) GenerateToken(ctx context.Context, userID string) (string, error) {
	return a.Jwtauth.GenerateToken(ctx, userID)
}

// ParseUserID returns a single string of userID
func (a *Auth) ParseUserID(ctx context.Context, tokenString string) (string, error) {
	return a.Jwtauth.ParseUserID(ctx, tokenString)
}
