package controllers

import (
	"MyWeb/models"
	"github.com/kataras/iris/mvc"
)

type OAuthController struct{}

func (m *OAuthController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/oauth/token", "Oauth")
}

func (m *OAuthController) Oauth(request models.ClientCredentialsTokenRequest) models.TokenResponse {

	return models.TokenResponse{AccessToken: "123", RefreshToken: request.ClientId + request.ClientSecret}
}
