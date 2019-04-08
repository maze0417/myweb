package controllers

import (
	"MyWeb/models"
	"MyWeb/services"
	"github.com/kataras/iris/mvc"
)

type PlayerController struct{}

func (m *PlayerController) BeforeActivation(b mvc.BeforeActivation) {
	//dep:=b.Dependencies()
	//dep.Add(func(ctx iris.Context) (request models.AuthorizePlayerRequest) {
	//	ctx.ReadJSON(&request)
	//	services.Validate(request)
	//	return
	//})

	b.Handle("POST", "/player/authorize", "Authorize")
}

func (m *PlayerController) Authorize(request models.AuthorizePlayerRequest, heloservice services.ITestService) models.AuthorizePlayerResponse {

	return models.AuthorizePlayerResponse{
		Err:     0,
		Errdesc: heloservice.MyTest(request.Userid),
	}
}
