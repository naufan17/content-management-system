package dtos

type LoginDto struct {
	Username string `json:"username" validate:"required,max=50"`
	Password string `json:"password" validate:"required,min=10"`
}

type AccessTokenDto struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
