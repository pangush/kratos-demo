package errcode

import (
	"github.com/bilibili/kratos/pkg/ecode"
)

//
var (
	// AuthPasswordErr .
	AuthPasswordErr = new(10000)
)

// new returns a single ecode.code
func new(code int) ecode.Code {
	return ecode.New(code)
}

// Register .
func Register() {
	cms := map[int]string{
		AuthPasswordErr.Code(): "账号或密码错误",
	}
	ecode.Register(cms)
}
