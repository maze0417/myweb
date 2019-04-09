package main

import (
	"MyWeb/controllers"
	"MyWeb/models"
	"MyWeb/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Use(OnException())
	requestLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(requestLogger)

	app.OnErrorCode(iris.StatusNotFound, notFound)

	mvc.Configure(app.Party("/api"), func(app *mvc.Application) {
		app.Register(
			//&services.Test2Service{},
			&services.TestService{},
			func(ctx iris.Context) (request models.AuthorizePlayerRequest) {
				ctx.ReadJSON(&request)
				services.Validate(request)
				return
			},
			func(ctx iris.Context) (request models.ClientCredentialsTokenRequest) {
				ctx.ReadForm(&request)
				services.Validate(request)
				return
			},
			)
		app.Handle(controllers.NewPlayerController())
		app.Handle(&controllers.OAuthController{})
	})

	app.Run(iris.Addr(":8080"), configure)
}

func notFound(ctx iris.Context) {
	ctx.JSON(models.CreateResponse(404, "Page not found"))
}

func configure(app *iris.Application) {
	app.Configure(
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
