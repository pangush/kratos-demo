package http

import (
	"kratos-demo/pkg/paginate"
	"kratos-demo/pkg/resp"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func testJSONWithPage(c *bm.Context) {
	var err error

	resp.JSONWithPage(c, nil, &paginate.Paginate{}, err)
}

func testJson(c *bm.Context) {
	var err error

	resp.JSON(c, nil, err)
}
