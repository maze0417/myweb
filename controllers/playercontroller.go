package controllers

import (
	"MyWeb/models"
	"MyWeb/services"
	"github.com/kataras/iris/mvc"
)

type PlayerController struct{}

var _helloService services.ITestService
func NewPlayerController(helloService services.ITestService) *PlayerController {
	_helloService=helloService
	return &PlayerController{}
}

func (m *PlayerController) BeforeActivation(b mvc.BeforeActivation) {
	//dep:=b.Dependencies()
	//dep.Add(func(ctx iris.Context) (request models.AuthorizePlayerRequest) {
	//	ctx.ReadJSON(&request)
	//	services.Validate(request)
	//	return
	//})

	b.Handle("POST", "/player/authorize", "Authorize")
}

func (m *PlayerController) Authorize(request models.AuthorizePlayerRequest) models.AuthorizePlayerResponse {

	return models.AuthorizePlayerResponse{
		Err:     0,
		Errdesc: _helloService.MyTest(request.Userid),
	}
}
