package bmcontext

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/metadata"
	"net/http"
	"strings"
)

// UserIDKey is a context key
const UserID = "kratosplugin_user_id"

// GetToken 获取用户令牌
func GetToken(c *bm.Context) string {
	var token string
	auth := c.Request.Header.Get("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// SetUserID 给bm context设置用户id
func SetUserID(ctx *bm.Context, userID string) {
	mid := userID
	ctx.Set(metadata.Mid, mid)
	if md, ok := metadata.FromContext(ctx); ok {
		md[metadata.Mid] = mid
		return
	}
}

// GetUserID 从bm context获取用户id
func GetUserID(ctx *bm.Context) (userID string) {
	mid, ok := ctx.Get(metadata.Mid)
	if ok {
		return mid.(string)
	}

	md, ok := metadata.FromContext(ctx)
	if ok {
		return md[metadata.Mid].(string)
	}

	return ""
}

//获取IP函数
func GetCurrentIP(r *http.Request) string {
	// 这里也可以通过X-Forwarded-For请求头的第一个值作为用户的ip
	// 但是要注意的是这两个请求头代表的ip都有可能是伪造的
	ip := r.Header.Get("X-Real-IP")
	if ip == ""{
		// 当请求头不存在即不存在代理时直接获取ip
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}
