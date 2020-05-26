package http

import (
	"kratos-demo/internal/conf"
	"kratos-demo/internal/server/http/middleware"
	"strings"

	"log"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func initRouter(e *bm.Engine) {
	e.Ping(ping)

	// 允许跨域
	corsDomains := strings.Split(conf.Conf.App.CorsDomains, ",")
	cors := bm.CORS(corsDomains)
	e.Use(cors)

	routePrefix := conf.Conf.App.AppName + "/v1"
	g := e.Group(routePrefix)

	g1 := g
	apiWithoutAuth(g1)

	jwtauth, err := middleware.NewAuth(conf.Conf.Jwt, conf.Conf.Redis)
	if err != nil {
		log.Fatalln(err)
	}

	g.Use(jwtauth.AuthMiddleware())

	api(g)
}

func api(api *bm.RouterGroup) {
	api.GET("/start", howToStart)
	api.GET("/demo", demo)

}

func apiWithoutAuth(api *bm.RouterGroup) {
	api.GET("/auth/login", authLogin)
	api.GET("/test/jsonwithpage", testJSONWithPage)
	api.GET("/test/json", testJson)
}
