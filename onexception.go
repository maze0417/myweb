// Package recover provides recovery for specific routes or for the whole app via middleware. See _examples/miscellaneous/recover
package main

import (
	"MyWeb/errors"
	"MyWeb/models"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"runtime"
	"strconv"
)

func getRequestLogs(ctx context.Context) string {
	var status, ip, method, path string
	status = strconv.Itoa(ctx.GetStatusCode())
	path = ctx.Path()
	method = ctx.Method()
	ip = ctx.RemoteAddr()
	// the date should be logged by iris' Logger, so we skip them
	return fmt.Sprintf("%v %s %s %s", status, path, method, ip)
}

func OnException() context.Handler {
	return func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() {
					return
				}

				var stacktrace string
				for i := 1; ; i++ {
					_, f, l, got := runtime.Caller(i)
					if !got {
						break

					}

					stacktrace += fmt.Sprintf("%s:%d\n", f, l)
				}

				// when stack finishes
				logMessage := fmt.Sprintf("Recovered from a route's Handler('%s')\n", ctx.HandlerName())
				logMessage += fmt.Sprintf("At Request: %s\n", getRequestLogs(ctx))
				logMessage += fmt.Sprintf("Trace: %s\n", err)
				logMessage += fmt.Sprintf("\n%s", stacktrace)
				ctx.Application().Logger().Warn(logMessage)

				if e, ok := err.(errors.ValidationError); ok {
					ctx.StatusCode(iris.StatusBadRequest)
					ctx.JSON(e)
					ctx.StopExecution()
					return
				}

				ctx.StatusCode(500)
				ctx.JSON(models.CreateResponse(900, fmt.Sprintf("%s", err)))
				ctx.StopExecution()
			}
		}()

		ctx.Next()
	}
}
