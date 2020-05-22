package input

type AuthLoginReq struct {
	Account string `json:"account"`
	Password string `json:"password"`
}

type AuthLoginResp struct {
	UserID int64 `json:"user_id"`
	Token string `json:"token"`
}
