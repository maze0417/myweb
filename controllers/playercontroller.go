package controllers

import (
	"MyWeb/models"
	"MyWeb/services"
	"github.com/kataras/iris/mvc"
)

type PlayerController struct{}

func (m *PlayerController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/player/authorize", "Authorize")
}

func (m *PlayerController) Authorize(request models.AuthorizePlayerRequest, helloService services.ITestService) models.AuthorizePlayerResponse {

	return models.AuthorizePlayerResponse{
		Err:     0,
		Errdesc: helloService.Hello(request.Userid),
	}
}
