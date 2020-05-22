package middleware

import (
	"fmt"
	"strings"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

// NoMethodHandler 未找到请求方法的处理函数
func NoMethodHandler() bm.HandlerFunc {
	return func(c *bm.Context) {
		c.AbortWithStatus(405)
	}
}

// NoRouteHandler 未找到请求路由的处理函数
func NoRouteHandler() bm.HandlerFunc {
	return func(c *bm.Context) {
		c.AbortWithStatus(404)
	}
}

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(*bm.Context) bool

// AllowPathPrefixSkipper 检查请求路径是否包含指定的前缀，如果包含则跳过
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *bm.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// AllowPathPrefixNoSkipper 检查请求路径是否包含指定的前缀，如果包含则不跳过
func AllowPathPrefixNoSkipper(prefixes ...string) SkipperFunc {
	return func(c *bm.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}
		return true
	}
}

// AllowMethodAndPathPrefixSkipper 检查请求方法和路径是否包含指定的前缀，如果不包含则跳过
func AllowMethodAndPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *bm.Context) bool {
		path := JoinRouter(c.Request.Method, c.Request.URL.Path)
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// JoinRouter 拼接路由
func JoinRouter(method, path string) string {
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(method), path)
}
