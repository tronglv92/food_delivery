package usermodel

type RenewAccessTokenRequest struct {
	Refreshtoken string `json:"refresh_token" form:"refresh_token" `
	AccessToken  string `json:"access_token" form:"access_token"`
}
