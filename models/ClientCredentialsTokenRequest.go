package models

type ClientCredentialsTokenRequest struct {
	ClientId     string `form:"client_id" validate:"required"`
	ClientSecret string `form:"client_secret" validate:"required"`
	GrantType    string `form:"grant_type"`
	Scope        string `form:"scope"`
}
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Err          int    `json:"err,omitempty"`
	Errdesc      string `json:"errdesc,omitempty"`
}
