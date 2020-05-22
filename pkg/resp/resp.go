package resp

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func JSON(c *bm.Context, data interface{}, err error) {
	if data == nil {
		c.JSON(map[string]interface{}{}, err)
		return
	}
	c.JSON(data, err)
	return
}

func AbortWithStatus(c *bm.Context, code int)  {

}