package resp

import (
	"kratos-demo/pkg/paginate"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

// JSON serializes the given struct as JSON into the response body.
// It also sets the Content-Type as "application/json".
func JSON(c *bm.Context, data interface{}, err error) {
	if data == nil {
		data = map[string]interface{}{}
	}
	c.JSON(data, err)
	return
}

// JSONWithPage serializes the given struct as JSON into the response body.
// And serializes the paginate.Paginate as JSON into the response body.
// It also sets the Content-Type as "application/json".
func JSONWithPage(c *bm.Context, data interface{}, page *paginate.Paginate, err error) {
	if data == nil {
		data = []interface{}{}
	}

	m := make(map[string]interface{})
	m["paginate"] = page
	m["list"] = data

	c.JSON(m, err)
	return
}
