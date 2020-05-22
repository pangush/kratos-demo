package service

import (
	"context"
	"kratos-demo/internal/server/http/input"
	"kratos-demo/internal/server/http/middleware"
	"strconv"
)

func (s *Service) AuthLogin(ctx context.Context, req *input.AuthLoginReq) (*input.AuthLoginResp, error) {
	data := &input.AuthLoginResp{}
	//todo 获取userID
	userID := int64(1)

	userIDStr := strconv.FormatInt(userID, 10)
	jwtToken, err := middleware.AuthMiddleware.GenerateToken(ctx, userIDStr)

	if err != nil {
		return data, err
	}

	data.UserID = userID
	data.Token = jwtToken

	return data, nil
}
