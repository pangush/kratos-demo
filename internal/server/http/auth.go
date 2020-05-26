package http

import (
	"kratos-demo/internal/server/http/input"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func authLogin(c *bm.Context) {
	params := &input.AuthLoginReq{}
	err := c.Bind(params)
	if err != nil {
		return
	}

	data, err := svc.AuthLogin(c.Context, params)

	c.JSON(data, err)
}
