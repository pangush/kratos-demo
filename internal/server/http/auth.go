package http

import (
	"kratos-demo/internal/errcode"
	"kratos-demo/internal/server/http/input"
	"kratos-demo/pkg/resp"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func authLogin(c *bm.Context) {
	params := &input.AuthLoginReq{}
	err := c.Bind(params)
	if err != nil {
		return
	}

	data, err := svc.AuthLogin(c.Context, params)
	resp.JSON(c, data, errcode.AuthPasswordErr)
	return
	c.JSON(data, err)
}
